package xyz.block.ftl.gradle

import com.squareup.kotlinpoet.AnnotationSpec
import com.squareup.kotlinpoet.ClassName
import com.squareup.kotlinpoet.FileSpec
import com.squareup.kotlinpoet.FunSpec
import com.squareup.kotlinpoet.KModifier
import com.squareup.kotlinpoet.ParameterSpec
import com.squareup.kotlinpoet.ParameterizedTypeName.Companion.parameterizedBy
import com.squareup.kotlinpoet.PropertySpec
import com.squareup.kotlinpoet.TypeName
import com.squareup.kotlinpoet.TypeSpec
import com.squareup.kotlinpoet.TypeVariableName
import org.gradle.configurationcache.extensions.capitalized
import xyz.block.ftl.Context
import xyz.block.ftl.Ignore
import xyz.block.ftl.Ingress
import xyz.block.ftl.v1.schema.Data
import xyz.block.ftl.v1.schema.Module
import xyz.block.ftl.v1.schema.Schema
import xyz.block.ftl.v1.schema.Type
import xyz.block.ftl.v1.schema.Verb
import java.io.File

class ModuleGenerator() {
  fun run(schema: Schema, outputDirectory: File, module: String) {
    schema.modules.filter { it.name != module }.forEach {
      val file = generateModule(it)
      file.writeTo(outputDirectory)

      println(
        "Generated module: ${outputDirectory.absolutePath}/ftl/${it.name}/${file.name}.kt"
      )
    }
  }

  internal fun generateModule(schema: Module): FileSpec {
    val namespace = "ftl.${schema.name}"
    val className = schema.name.capitalized()
    val file = FileSpec.builder(namespace, className)
      .addFileComment("Code generated by FTL-Plugin, do not edit.")

    schema.comments?.let {
      file.addFileComment("\n")
      file.addFileComment(it.joinToString("\n"))
    }

    val moduleClass = TypeSpec.classBuilder(className)
      .addAnnotation(AnnotationSpec.builder(Ignore::class).build())
      .primaryConstructor(
        FunSpec.constructorBuilder().build()
      )

    val types = schema.decls?.mapNotNull { it.data_ } ?: listOf()
    types.forEach { file.addType(buildDataClass(it)) }

    val verbs = schema.decls?.mapNotNull { it.verb } ?: listOf()
    verbs.forEach { moduleClass.addFunction(buildVerbFunction(className, it)) }

    file.addType(moduleClass.build())
    return file.build()
  }

  private fun buildDataClass(type: Data): TypeSpec {
    val dataClassBuilder = TypeSpec.classBuilder(type.name)
      .addModifiers(KModifier.DATA)
      .addKdoc(type.comments.joinToString("\n"))

    val dataConstructorBuilder = FunSpec.constructorBuilder()
    type.fields.forEach { field ->
      dataClassBuilder.addKdoc(field.comments.joinToString("\n"))
      field.type?.let {
        dataConstructorBuilder.addParameter(field.name, getTypeClass(it))
        dataClassBuilder.addProperty(
          PropertySpec.builder(field.name, getTypeClass(it)).initializer(field.name).build()
        )
      }
    }

    // Handle empty data classes.
    if (type.fields.isEmpty()) {
      dataConstructorBuilder.addParameter(
        ParameterSpec.builder("_empty", Unit::class).defaultValue("Unit").build()
      )
      dataClassBuilder.addProperty(
        PropertySpec.builder("_empty", Unit::class).initializer("_empty").build()
      )
    }

    dataClassBuilder.primaryConstructor(dataConstructorBuilder.build())

    return dataClassBuilder.build()
  }

  private fun buildVerbFunction(className: String, verb: Verb): FunSpec {
    val verbFunBuilder =
      FunSpec.builder(verb.name).addKdoc(verb.comments.joinToString("\n")).addAnnotation(
        AnnotationSpec.builder(xyz.block.ftl.Verb::class).build()
      )

    verb.metadata.forEach { metadata ->
      metadata.ingress?.let {
        verbFunBuilder.addAnnotation(
          AnnotationSpec.builder(Ingress::class)
            .addMember("%T", ClassName("xyz.block.ftl.Method", it.method))
            .addMember("%S", it.path)
            .build()
        )
      }
    }

    verbFunBuilder.addParameter("context", Context::class)

    verb.request?.let {
      verbFunBuilder.addParameter(
        "req", TypeVariableName(it.name)
      )
    }

    verb.response?.let {
      verbFunBuilder.returns(TypeVariableName(it.name))
    }

    val message =
      "Verb stubs should not be called directly, instead use context.call($className::${verb.name}, ...)"
    verbFunBuilder.addCode("""throw NotImplementedError(%S)""", message)

    return verbFunBuilder.build()
  }

  private fun getTypeClass(type: Type): TypeName {
    return when {
      type.int != null -> ClassName("kotlin", "Long")
      type.float != null -> ClassName("kotlin", "Float")
      type.string != null -> ClassName("kotlin", "String")
      type.bool != null -> ClassName("kotlin", "Boolean")
      type.time != null -> ClassName("java.time", "OffsetDateTime")
      type.array != null -> {
        val element = type.array.element ?: throw IllegalArgumentException(
          "Missing element type in kotlin array generator"
        )
        val elementType = getTypeClass(element)
        val arrayList = ClassName("kotlin.collections", "ArrayList")
        arrayList.parameterizedBy(elementType)
      }

      type.map != null -> {
        val map = ClassName("kotlin.collections", "Map")
        val key =
          type.map.key ?: throw IllegalArgumentException("Missing map key in kotlin map generator")
        val value = type.map.value_ ?: throw IllegalArgumentException(
          "Missing map value in kotlin map generator"
        )
        map.parameterizedBy(getTypeClass(key), getTypeClass(value))
      }

      type.verbRef != null -> ClassName("xyz.block.ftl.v1.schema", "VerbRef")
      type.dataRef != null -> ClassName("xyz.block.ftl.v1.schema", "DataRef")

      else -> throw IllegalArgumentException("Unknown type in kotlin generator")
    }
  }
}
