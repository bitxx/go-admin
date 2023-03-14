<template>
  <el-dialog :close-on-click-modal="false" title="选择文章分类" :visible.sync="open" width="800px" top="5vh" append-to-body>
    <el-form ref="queryForm" :model="queryParams" :inline="true" label-width="68px">
      <el-form-item label-width="100" label="分类编号" prop="id">
        <el-input v-model="queryParams.id" placeholder="请输入分类编号" clearable size="small" @keyup.enter.native="handleQuery" />
      </el-form-item>
      <el-form-item label-width="100" label="分类名称" prop="name">
        <el-input v-model="queryParams.name" placeholder="请输入分类名称" clearable size="small" @keyup.enter.native="handleQuery" />
      </el-form-item>

      <el-form-item>
        <el-button type="primary" icon="el-icon-search" size="mini" @click="handleQuery">搜索</el-button>
        <el-button icon="el-icon-refresh" size="mini" @click="resetQuery">重置</el-button>
      </el-form-item>
    </el-form>

    <el-table v-loading="loading" stripe border :data="categoryList" height="260" @row-click="clickRow">
      <el-table-column label="序号" type="index" align="center" width="80">
        <template slot-scope="scope">
          <span>{{ (queryParams.pageIndex - 1) * queryParams.pageSize + scope.$index + 1 }}</span>
        </template>
      </el-table-column>
      <el-table-column width="100" label="分类编号" align="center" prop="id" :show-overflow-tooltip="true" />
      <el-table-column width="200" label="分类名称" align="center" prop="name" :show-overflow-tooltip="true" />
      <el-table-column width="100" label="更新人编号" align="center" prop="updateBy" :show-overflow-tooltip="true" />
      <el-table-column label="创建时间" align="center" prop="createdAt" :show-overflow-tooltip="true">
        <template slot-scope="scope">
          <span>{{ parseTime(scope.row.createdAt) }}</span>
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
  </el-dialog>
</template>

<script>

import { listContentCategory } from '@/api/plugins/content/content-category'

export default {
  data() {
    return {
      // 遮罩层
      loading: true,
      // 总条数
      total: 0,
      // 弹出层标题
      title: '',
      // 是否显示弹出层
      open: false,
      // 类型数据字典
      categoryList: [],

      // 查询参数
      queryParams: {
        pageIndex: 1,
        pageSize: 10,
        name: undefined
      },
      // 表单参数
      form: {}
    }
  },
  created() {
    this.getList()
  },
  methods: {
    // 显示弹框
    show() {
      this.queryParams = {
        pageIndex: 1,
        pageSize: 20,
        name: undefined
      }
      this.getList()
      this.open = true
    },
    /** 查询参数列表 */
    getList() {
      this.loading = true
      listContentCategory(this.addDateRange(this.queryParams, this.dateRange))
        .then(response => {
          this.categoryList = response.data.list
          this.total = response.data.count
          this.loading = false
        })
    },
    clickRow(row) {
      this.open = false
      this.$emit('ok', row)
    },
    /** 搜索按钮操作 */
    handleQuery() {
      this.queryParams.pageIndex = 1
      this.getList()
    },
    /** 重置按钮操作 */
    resetQuery() {
      this.dateRange = []
      this.resetForm('queryForm')
      this.handleQuery()
    }
  }
}
</script>
