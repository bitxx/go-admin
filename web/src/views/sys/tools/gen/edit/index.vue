<template>
  <el-card>
    <el-tabs v-model="activeName">
      <el-tab-pane label="基本信息" name="basic">
        <basic-info-form ref="basicInfo" :info="info" />
      </el-tab-pane>
      <el-tab-pane label="字段信息" name="cloum">
        <el-alert title="⚠️还没想好要警告什么~" type="warning" show-icon />
        <el-table :data="info.sysGenColumns" :max-height="tableHeight" stripe border style="width: 100%">
          <el-table-column fixed label="序号" type="index" width="50" align="center" />
          <el-table-column fixed label="字段列名" prop="columnName" width="180" align="center" :show-overflow-tooltip="true" />
          <el-table-column fixed label="字段描述" width="180" align="center">
            <template slot-scope="scope">
              <el-input v-model="scope.row.columnComment" />
            </template>
          </el-table-column>
          <el-table-column label="物理类型" prop="columnType" align="center" width="120" :show-overflow-tooltip="true" />
          <el-table-column label="go类型" width="120" align="center">
            <template slot-scope="scope">
              <el-select v-model="scope.row.goType">
                <el-option
                  v-for="dict in sysGenGoTypeOptions"
                  :key="dict.dictValue"
                  :label="dict.dictLabel"
                  :value="dict.dictValue"
                />
              </el-select>
            </template>
          </el-table-column>
          <el-table-column label="go属性" width="150" align="center">
            <template slot-scope="scope">
              <el-input v-model="scope.row.goField" />
            </template>
          </el-table-column>
          <el-table-column label="json属性" width="150" align="center">
            <template slot-scope="scope">
              <el-input v-model="scope.row.jsonField" />
            </template>
          </el-table-column>
          <el-table-column label="查询方式" width="120" align="center">
            <template slot-scope="scope">
              <el-select v-model="scope.row.queryType">
                <el-option
                  v-for="dict in sysGenQueryTypeOptions"
                  :key="dict.dictValue"
                  :label="dict.dictLabel"
                  :value="dict.dictValue"
                />
              </el-select>
            </template>
          </el-table-column>
          <el-table-column label="显示类型" width="140" align="center">
            <template slot-scope="scope">
              <el-select v-model="scope.row.htmlType">
                <el-option
                  v-for="dict in sysGenHtmlTypeOptions"
                  :key="dict.dictValue"
                  :label="dict.dictLabel"
                  :value="dict.dictValue"
                />
              </el-select>
            </template>
          </el-table-column>
          <el-table-column label="字典类型" width="200" align="center">
            <template slot-scope="scope">
              <el-select v-model="scope.row.dictType" clearable filterable placeholder="请选择">
                <el-option
                  v-for="dict in dictOptions"
                  :key="dict.dictType"
                  :label="dict.dictName"
                  :value="dict.dictType"
                >
                  <span style="float: left">{{ dict.dictName }}</span>
                  <span style="float: right; color: #8492a6; font-size: 13px">{{ dict.dictType }}</span>
                </el-option>
              </el-select>
            </template>
          </el-table-column>
          <el-table-column align="center" label="编辑" width="70">
            <template slot-scope="scope">
              <el-checkbox v-model="scope.row.isRequired" true-label="1" false-label="2" />
            </template>
          </el-table-column>
          <el-table-column align="center" label="列表" width="70" :render-header="renderHeadeList" :cell-style="{'text-align':'center'}">
            <template slot-scope="scope">
              <el-checkbox v-model="scope.row.isList" true-label="1" false-label="0" />
            </template>
          </el-table-column>
          <el-table-column label="查询" align="center" width="70" :render-header="renderHeadeSearch" :cell-style="{'text-align':'center'}">
            <template slot-scope="scope">
              <el-checkbox v-model="scope.row.isQuery" true-label="1" false-label="2" />
            </template>
          </el-table-column>
        </el-table>
      </el-tab-pane>
      <el-tab-pane label="生成信息" name="genInfo">
        <gen-info-form ref="genInfo" :info="info" />
      </el-tab-pane>
    </el-tabs>
    <el-form label-width="100px">
      <el-form-item style="text-align: center;margin-left:-100px;margin-top:10px;">
        <el-button type="primary" @click="submitForm()">提交</el-button>
        <el-button @click="close()">返回</el-button>
      </el-form-item>
    </el-form>
  </el-card>
</template>
<script>
import { getGenTable, updateGenTable } from '@/api/sys/tools/table'
import { optionselect as getDictOptionselect } from '@/api/sys/dicttype'
import basicInfoForm from '../basicInfoForm.vue'
import genInfoForm from '../genInfoForm.vue'
export default {
  name: 'SysEditTable',
  components: {
    basicInfoForm,
    genInfoForm
  },
  data() {
    return {
      // 选中选项卡的 name
      activeName: 'cloum',
      // 表格的高度
      tableHeight: document.documentElement.scrollHeight - 245 + 'px',
      // 字典信息
      dictOptions: [],
      sysGenGoTypeOptions: [],
      sysGenQueryTypeOptions: [],
      sysGenHtmlTypeOptions: [],
      // 表详细信息
      info: {}
    }
  },

  beforeCreate() {
    const { tableId } = this.$route.query
    if (tableId) {
      // 获取表详细信息
      getGenTable(tableId).then(res => {
        this.info = res.data
      })

      /** 查询字典下拉列表 */
      getDictOptionselect().then(response => {
        this.dictOptions = response.data
      })
      this.getDicts('sys_gen_go_type').then(response => {
        this.sysGenGoTypeOptions = response.data
      })
      this.getDicts('sys_gen_query_type').then(response => {
        this.sysGenQueryTypeOptions = response.data
      })
      this.getDicts('sys_gen_html_type').then(response => {
        this.sysGenHtmlTypeOptions = response.data
      })
    }
  },
  methods: {
    renderHeadeSearch(h, { column, $index }) {
      return h('div', [
        h('span', column.label + '  ', { align: 'center', marginTop: '0px' }),
        h(
          'el-popover',
          { props: { placement: 'top-start', width: '270', trigger: 'hover' }},
          [
            h('p', '是都当做搜索条件，打√表示做为搜索条件', { class: 'text-align: center; margin: 0' }),
            h('i', { class: 'el-icon-question', style: 'color:#ccc,padding-top:5px', slot: 'reference' })
          ]
        )
      ])
    },
    renderHeadeList(h, { column, $index }) {
      // h 是一个渲染函数       column 是一个对象表示当前列      $index 第几列
      return h('div', [
        h('span', column.label + '  ', { align: 'center', marginTop: '0px' }),
        h(
          'el-popover',
          { props: { placement: 'top-start', width: '260', trigger: 'hover' }},
          [
            h('p', '是否在列表中展示，打√表示需要展示', { class: 'text-align: center; margin: 0' }),
            h('i', { class: 'el-icon-question', style: 'color:#ccc,padding-top:5px', slot: 'reference' })
          ]
        )
      ])
    },
    /** 提交按钮 */
    submitForm() {
      const basicForm = this.$refs.basicInfo.$refs.basicInfoForm
      const genForm = this.$refs.genInfo.$refs.genInfoForm

      Promise.all([basicForm, genForm].map(this.getFormPromise)).then(res => {
        const validateResult = res.every(item => !!item)
        if (validateResult) {
          updateGenTable(this.info, this.info.id).then(res => {
            this.msgSuccess(res.msg)
            if (res.code === 200) {
              this.close()
            }
          })
        } else {
          this.msgError('表单校验未通过，请重新检查提交内容')
        }
      })
    },
    getFormPromise(form) {
      return new Promise(resolve => {
        form.validate(res => {
          resolve(res)
        })
      })
    },
    /** 关闭按钮 */
    close() {
      this.$store.dispatch('tagsView/delView', this.$route)
      this.$router.push({ path: '/sys-tools/sys-gen', query: { t: Date.now() }})
    }
  }
}
</script>
