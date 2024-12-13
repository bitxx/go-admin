<template>
  <el-form ref="genInfoForm" :model="info" :rules="rules" label-width="150px">
    <el-row>
      <el-col :span="12">
        <el-form-item prop="packageName">
          <span slot="label">
            应用名
            <el-tooltip content="应用名，建议使用纯小写单词，默认是表名前缀，如app_user_account，则取app" placement="top">
              <i class="el-icon-question" />
            </el-tooltip>
          </span>
          <el-input v-model="info.packageName" />
        </el-form-item>
      </el-col>
      <el-col :span="12">
        <el-form-item prop="businessName">
          <span slot="label">
            业务名
            <el-tooltip content="默认表名第二部分，建议使用小写单词和下划线组成，如app_user_account，则取user" placement="top">
              <i class="el-icon-question" />
            </el-tooltip>
          </span>
          <el-input v-model="info.businessName" />
        </el-form-item>
      </el-col>

      <el-col :span="12">
        <el-form-item prop="functionName">
          <span slot="label">
            功能描述
            <el-tooltip content="同步的数据库表备注，用作类描述，例如：用户" placement="top">
              <i class="el-icon-question" />
            </el-tooltip>
          </span>
          <el-input v-model="info.functionName" />
        </el-form-item>
      </el-col>
      <el-col :span="12">
        <el-form-item prop="moduleName">
          <span slot="label">
            接口路径
            <el-tooltip content="接口路径，建议使用小写单词和下划线组成，默认取表名除前缀的内容(下划线替换为中划线)，如app_user_account，则取user-account" placement="top">
              <i class="el-icon-question" />
            </el-tooltip>
          </span>
          <el-input v-model="info.moduleName">
            <template slot="prepend">api/{version}/</template>
            <template slot="append">...</template>
          </el-input>
        </el-form-item>
      </el-col>
    </el-row>
  </el-form>
</template>
<script>
export default {
  name: 'BasicInfoForm',
  props: {
    info: {
      type: Object,
      default: null
    }
  },
  data() {
    return {
      // 状态数据字典
      rules: {
        packageName: [
          { required: true, message: '请输入生成包路径', trigger: 'blur' }
          // { pattern: /^[a-z]*$/g, trigger: 'blur', message: '只允许小写字母,例如 system 格式' }
        ],
        moduleName: [
          { required: true, message: '请输入生成模块名', trigger: 'blur' }
          // { pattern: /^[a-z\-]*[a-z]$/g, trigger: 'blur', message: '只允许小写字母,例如 sys-demo 格式' }
        ],
        businessName: [
          { required: true, message: '请输入生成业务名', trigger: 'blur' }
          // { pattern: /^[a-z][A-Za-z]+$/, trigger: 'blur', message: '校验规则:  只允许输入字母 a-z 或大写 A-Z ，并且小写字母开头' }
        ],
        functionName: [
          { required: true, message: '请输入生成功能名', trigger: 'blur' }
        ]
      }
    }
  },
  created() {},
  methods: {}
}
</script>
