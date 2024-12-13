<template>
  <BasicLayout>
    <template #wrapper>
      <el-card class="box-card">
        <el-form ref="queryForm" :model="queryParams" :inline="true" label-position="left">
          <el-form-item label-width="100" label="数据库表名" prop="tableName">
            <el-input v-model="queryParams.tableName" placeholder="请输入数据库表名" clearable size="small" @keyup.enter.native="handleQuery" />
          </el-form-item>
          <el-form-item label-width="100" label="数据库表描述" prop="tableComment">
            <el-input v-model="queryParams.tableComment" placeholder="请输入数据库表描述" clearable size="small" @keyup.enter.native="handleQuery" />
          </el-form-item>
          <el-form-item label="创建时间">
            <el-date-picker
              v-model="dateRange"
              size="small"
              type="datetimerange"
              range-separator="至"
              start-placeholder="开始日期"
              end-placeholder="结束日期"
              align="right"
              value-format="yyyy-MM-dd HH:mm:ss"
            />
          </el-form-item>
          <el-form-item>
            <el-button type="primary" icon="el-icon-search" size="mini" @click="handleQuery">搜索</el-button>
            <el-button icon="el-icon-refresh" size="mini" @click="resetQuery">重置</el-button>
          </el-form-item>
        </el-form>
        <el-row :gutter="10" class="mb8">
          <el-col :span="1.5">
            <el-button
              v-permisaction="['sys:table:add']"
              type="warning"
              icon="el-icon-upload"
              size="mini"
              @click="openImportTable"
            >导入数据库表</el-button>
          </el-col>
        </el-row>
        <el-table v-loading="loading" stripe border :data="tableList">
          <el-table-column label="序号" type="index" align="center" width="80">
            <template slot-scope="scope">
              <span>{{ (queryParams.pageIndex - 1) * queryParams.pageSize + scope.$index + 1 }}</span>
            </template>
          </el-table-column>
          <el-table-column width="80" label="表编号" align="center" prop="id" :show-overflow-tooltip="true" />
          <el-table-column width="220" label="数据库表名" align="center" prop="tableName" :show-overflow-tooltip="true" />
          <el-table-column width="150" label="数据库表描述" align="center" prop="tableComment" :show-overflow-tooltip="true" />
          <el-table-column width="200" label="类名" align="center" prop="className" :show-overflow-tooltip="true" />
          <el-table-column width="180" label="创建时间" align="center" prop="createdAt" :show-overflow-tooltip="true">
            <template slot-scope="scope">
              <span>{{ parseTime(scope.row.createdAt) }}</span>
            </template>
          </el-table-column>
          <el-table-column label="操作" align="center" class-name="small-padding fixed-width">
            <template slot-scope="scope">
              <el-button
                v-permisaction="['admin:sys-gen:edit']"
                type="text"
                size="small"
                icon="el-icon-edit"
                @click="handleEditTable(scope.row)"
              >编辑</el-button>
              <el-button
                type="text"
                size="small"
                icon="el-icon-view"
                @click="handlePreview(scope.row)"
              >预览</el-button>
              <el-button
                size="mini"
                type="text"
                icon="el-icon-delete"
                @click="handleGenCode(scope.row)"
              >生成代码
              </el-button>
              <el-button
                size="mini"
                type="text"
                icon="el-icon-delete"
                @click="handleDownloadCode(scope.row)"
              >下载代码
              </el-button>
              <el-button
                size="mini"
                type="text"
                icon="el-icon-delete"
                @click="handleGenDB(scope.row)"
              >配置生成
              </el-button>
              <el-button
                size="mini"
                type="text"
                icon="el-icon-delete"
                @click="handleDelete(scope.row)"
              >删除
              </el-button>
            </template>
          </el-table-column>
        </el-table>
        <pagination
          v-show="total>0"
          :total="total"
          :page.sync="queryParams.pageIndex"
          :limit.sync="queryParams.pageSize"
          @pagination="getList"
        />
      </el-card>

      <!-- 预览界面 -->

      <el-dialog class="preview" :title="preview.title" :visible.sync="preview.open" fullscreen>
        <div class="el-dialog-container">
          <div class="tag-group">
            <!-- eslint-disable-next-line vue/valid-v-for -->
            <el-tag v-for="item in preview.data" :key="item.id" @click="codeChange(item)">
              <template>
                {{ item.name }}
              </template>
            </el-tag>
          </div>
          <div id="codemirror">
            <codemirror ref="cmEditor" :value="codestr" :options="cmOptions" />
          </div>
        </div>

      </el-dialog>
      <import-table ref="importTB" @ok="handleQuery" />
    </template>
  </BasicLayout>
</template>

<script>
import { listTable, previewTable, delTable, genCode, downloadCode, genDB } from '@/api/admin/sys/tools/table'
import importTable from './importTable.vue'
import { codemirror } from 'vue-codemirror'
import 'codemirror/theme/material-palenight.css'

require('codemirror/mode/javascript/javascript')

import 'codemirror/mode/javascript/javascript'
import 'codemirror/mode/go/go'
import 'codemirror/mode/vue/vue'
import { resolveBlobNoHeader } from '@/utils/zipdownload'

export default {
  name: 'SysGen',
  components: { importTable, codemirror },
  data() {
    return {
      cmOptions: {
        tabSize: 4,
        theme: 'material-palenight',
        mode: 'text/javascript',
        lineNumbers: true,
        line: true
        // more CodeMirror options...
      },
      codestr: '',
      // 遮罩层
      loading: true,
      // 选中表数组
      tableNames: [],
      // 总条数
      total: 0,
      // 表数据
      tableList: [],
      // 日期范围
      dateRange: '',
      // 查询参数
      queryParams: {
        pageIndex: 1,
        pageSize: 10,
        tableName: undefined,
        tableComment: undefined
      },
      // 预览参数
      preview: {
        open: false,
        title: '代码预览',
        data: {}
      }
    }
  },
  created() {
    this.getList()
  },
  methods: {
    /** 查询表集合 */
    getList() {
      this.loading = true
      listTable(this.addDateRange(this.queryParams, this.dateRange)).then(response => {
        this.tableList = response.data.list
        this.total = response.data.count
        this.loading = false
      })
    },
    codeChange(item) {
      if (item.name.indexOf('js') > -1) {
        this.cmOptions.mode = 'text/javascript'
      }
      if (item.name.indexOf('model') > -1 || item.name.indexOf('router') > -1 || item.name.indexOf('businessRouter') > -1 || item.name.indexOf('api') > -1 || item.name.indexOf('constant') > -1 || item.name.indexOf('service') > -1 || item.name.indexOf('dto') > -1) {
        this.cmOptions.mode = 'text/x-go'
      }
      if (item.name.indexOf('vue') > -1) {
        this.cmOptions.mode = 'text/x-vue'
      }
      this.codestr = item.content
    },
    /** 搜索按钮操作 */
    handleQuery() {
      this.queryParams.pageIndex = 1
      this.getList()
    },
    /** 打开导入表弹窗 */
    openImportTable() {
      this.$refs.importTB.show()
    },
    /** 重置按钮操作 */
    resetQuery() {
      this.dateRange = []
      this.resetForm('queryForm')
      this.handleQuery()
    },
    /** 预览按钮 */
    handlePreview(row) {
      previewTable(row.id).then(response => {
        this.preview.data = response.data
        this.preview.open = true
        this.codeChange(this.preview.data[0])
      })
    },

    handleGenCode(row) {
      this.open = true
      this.$confirm('是否为编号"' + row.id + '"生成代码？', '警告', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(function() {
        return genCode(row.id)
      }).then((response) => {
        this.open = false
        if (response.code === 200) {
          this.msgSuccess(response.msg)
        } else {
          this.msgError(response.msg)
        }
      }).catch(function() {
      })
    },
    handleDownloadCode(row) {
      this.open = true
      this.$confirm('是否为编号"' + row.id + '"下载代码？', '警告', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(function() {
        return downloadCode(row.id)
      }).then((response) => {
        this.open = false
        resolveBlobNoHeader(response, 'code.zip')
      }).catch(function() {
      })
    },

    handleGenDB(row) {
      this.open = true
      this.$confirm('是否为编号"' + row.id + '"导入菜单到数据库？', '警告', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(function() {
        return genDB(row.id)
      }).then((response) => {
        this.open = false
        if (response.code === 200) {
          this.msgSuccess(response.msg)
        } else {
          this.msgError(response.msg)
        }
      }).catch(function() {
      })
    },
    /** 修改按钮操作 */
    handleEditTable(row) {
      this.$router.push({ path: '/admin/sys/tools/sys-edit-table', query: { tableId: row.id }})
    },
    /** 删除按钮操作 */
    handleDelete(row) {
      this.open = true
      const ids = [row.id]
      this.$confirm('是否确认删除编号为"' + ids + '"的数据项?', '警告', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(function() {
        return delTable({ 'ids': ids })
      }).then((response) => {
        this.open = false
        if (response.code === 200) {
          this.msgSuccess(response.msg)
          this.getList()
        } else {
          this.msgError(response.msg)
        }
      }).catch(function() {
      })
    }
  }
}
</script>

<style lang="scss" scoped>
 .el-dialog-container ::v-deep{
   height:600px;
   overflow: hidden;
   .el-scrollbar__view{
     height: 100%;
   }
   .pre{
     height: 546px;
      overflow: hidden;
      .el-scrollbar{
        height: 100%;
      }
   }
   .el-scrollbar__wrap::-webkit-scrollbar{
     display: none;
   }
 }
 ::v-deep .el-dialog__body{
    padding: 0 20px;
    margin:0;
  }

   .tag-group .el-tag{
    margin-left: 10px;
  }

</style>

<style lang="scss">
  #codemirror {
      height: auto;
      margin: 0;
      overflow: auto;
    }
  .CodeMirror {
      overflow:auto;
      border: 1px solid #eee;
      height: 600px;
    }
</style>
