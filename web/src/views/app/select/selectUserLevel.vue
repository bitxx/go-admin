<template>
  <el-dialog :close-on-click-modal="false" title="选择等级" :visible.sync="open" width="800px" top="5vh" append-to-body>
    <el-form ref="queryForm" :model="queryParams" :inline="true" label-width="68px">
      <el-form-item label-width="100" label="等级名称" prop="name">
        <el-input v-model="queryParams.name" placeholder="请输入等级名称" clearable size="small" @keyup.enter.native="handleQuery" />
      </el-form-item>
      <el-form-item label-width="100" label="等级类型" prop="levelType">
        <el-select v-model="queryParams.levelType" placeholder="用户等级等级类型" clearable size="small">
          <el-option
            v-for="dict in levelTypeOptions"
            :key="dict.dictValue"
            :label="dict.dictLabel"
            :value="dict.dictValue"
          />
        </el-select>
      </el-form-item>
      <el-form-item label="等级" prop="level">
        <el-input v-model="queryParams.level" placeholder="请输入等级" clearable size="small" @keyup.enter.native="handleQuery" />
      </el-form-item>

      <el-form-item>
        <el-button type="primary" icon="el-icon-search" size="mini" @click="handleQuery">搜索</el-button>
        <el-button icon="el-icon-refresh" size="mini" @click="resetQuery">重置</el-button>
      </el-form-item>
    </el-form>

    <el-table v-loading="loading" stripe border :data="tableList" height="260px" @row-click="clickRow">
      <el-table-column label="序号" type="index" align="center" width="80">
        <template slot-scope="scope">
          <span>{{ (queryParams.pageIndex - 1) * queryParams.pageSize + scope.$index + 1 }}</span>
        </template>
      </el-table-column>
      <el-table-column label="等级编号" align="center" prop="id" :show-overflow-tooltip="true" />
      <el-table-column label="等级名称" align="center" prop="name" :show-overflow-tooltip="true" />
      <el-table-column label="等级类型" align="center" prop="levelType" :formatter="levelTypeFormat" width="100">
        <template slot-scope="scope">
          {{ levelTypeFormat(scope.row) }}
        </template>
      </el-table-column>
      <el-table-column label="等级" align="center" prop="level" :show-overflow-tooltip="true" />
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
import { listUserLevel } from '@/api/app/user/user-level'

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
      tableList: [],
      levelTypeOptions: [],

      // 查询参数
      queryParams: {
        pageIndex: 1,
        pageSize: 20,
        id: undefined,
        name: undefined,
        levelType: undefined,
        level: undefined
      },
      // 表单参数
      form: {}
    }
  },
  created() {
    this.getList()
    this.getDicts('app_user_level_type').then(response => {
      this.levelTypeOptions = response.data
    })
  },
  methods: {
    // 显示弹框
    show(levelType) {
      this.queryParams = {
        id: undefined,
        pageIndex: 1,
        pageSize: 20,
        name: undefined,
        levelType: levelType,
        level: undefined
      }
      this.getList()
      this.open = true
    },
    /** 查询参数列表 */
    getList() {
      this.loading = true
      listUserLevel(this.addDateRange(this.queryParams, this.dateRange)).then(response => {
        this.tableList = response.data.list
        this.total = response.data.count
        this.loading = false
      })
    },
    clickRow(row) {
      this.open = false
      this.$emit('ok', row)
    },
    levelTypeFormat(row) {
      return this.selectDictLabel(this.levelTypeOptions, row.levelType)
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
