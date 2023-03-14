<template>
  <BasicLayout>
    <template #wrapper>
      <el-card class="box-card">
        <el-form ref="queryForm" :model="queryParams" :inline="true" label-width="68px">
          <el-form-item label="用户编号" prop="userId">
            <el-input v-model="queryParams.userId" placeholder="请输入用户编号" clearable size="small" @keyup.enter.native="handleQuery" />
          </el-form-item>
          <el-form-item label-width="100" label="用户名" prop="userName">
            <el-input v-model="queryParams.userName" placeholder="请输入昵称" clearable size="small" @keyup.enter.native="handleQuery" />
          </el-form-item>
          <el-form-item label-width="100" label="用户邮箱" prop="email">
            <el-input v-model="queryParams.email" placeholder="请输入用户邮箱" clearable size="small" @keyup.enter.native="handleQuery" />
          </el-form-item>
          <el-form-item label-width="100" label="用户手机号" prop="mobile">
            <el-input v-model="queryParams.mobile" placeholder="请输入用户手机号" clearable size="small" @keyup.enter.native="handleQuery" />
          </el-form-item>
          <el-form-item label-width="120" label="用户行为类型" prop="actionType">
            <el-select v-model="queryParams.actionType" placeholder="用户行为类型" clearable size="small">
              <el-option
                v-for="dict in actionTypeOptions"
                :key="dict.dictValue"
                :label="dict.dictLabel"
                :value="dict.dictValue"
              />
            </el-select>
          </el-form-item>
          <el-form-item label-width="120" label="更新用户类型" prop="byType">
            <el-select v-model="queryParams.byType" placeholder="更新用户类型" clearable size="small">
              <el-option
                v-for="dict in byTypeOptions"
                :key="dict.dictValue"
                :label="dict.dictLabel"
                :value="dict.dictValue"
              />
            </el-select>
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
              v-permisaction="['app:user:user-oper-log:export']"
              type="success"
              icon="el-icon-download"
              size="mini"
              @click="handleExport"
            >
              Excel导出
            </el-button>
          </el-col>
        </el-row>

        <el-table v-loading="loading" stripe border :data="tableList">
          <el-table-column label="序号" type="index" align="center" width="80">
            <template slot-scope="scope">
              <span>{{ (queryParams.pageIndex - 1) * queryParams.pageSize + scope.$index + 1 }}</span>
            </template>
          </el-table-column>
          <el-table-column label="日志编号" align="center" prop="id" :show-overflow-tooltip="true" />
          <el-table-column label="用户编号" align="center" prop="userId" :show-overflow-tooltip="true" />
          <el-table-column width="150" label="昵称" align="center" prop="user.userName" :show-overflow-tooltip="true" />
          <el-table-column width="150" label="用户邮箱" align="center" prop="user.email" :show-overflow-tooltip="true" />
          <el-table-column width="150" label="用户手机号" align="center" prop="user.mobile" :show-overflow-tooltip="true" />
          <el-table-column label="用户行为类型" align="center" prop="actionType" :formatter="actionTypeFormat" width="150">
            <template slot-scope="scope">
              {{ actionTypeFormat(scope.row) }}
            </template>
          </el-table-column>
          <el-table-column label="更新用户类型" align="center" prop="byType" :formatter="byTypeFormat" width="100">
            <template slot-scope="scope">
              {{ byTypeFormat(scope.row) }}
            </template>
          </el-table-column>
          <el-table-column label="创建者" align="center" prop="createBy" :show-overflow-tooltip="true" />
          <el-table-column label="更新者" align="center" prop="updateBy" :show-overflow-tooltip="true" />
          <el-table-column width="200" label="创建时间" align="center" prop="createdAt" :show-overflow-tooltip="true">
            <template slot-scope="scope">
              <span>{{ parseTime(scope.row.createdAt) }}</span>
            </template>
          </el-table-column>
          <el-table-column width="200" label="更新时间" align="center" prop="updatedAt" :show-overflow-tooltip="true">
            <template slot-scope="scope">
              <span>{{ parseTime(scope.row.updatedAt) }}</span>
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
    </template>
  </BasicLayout>
</template>

<script>
import { listUserOperLog, exportUserOperLog } from '@/api/app/user/user-oper-log'
import { resolveBlob } from '@/utils/download'
export default {
  name: 'UserOperLog',
  components: {},
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
      isEdit: false,
      // 日期范围
      dateRange: [],
      // 类型数据字典
      tableList: [],
      actionTypeOptions: [],
      byTypeOptions: [],
      // 查询参数
      queryParams: {
        pageIndex: 1,
        pageSize: 10,
        actionType: undefined,
        byType: undefined,
        userId: undefined
      },
      // 表单参数
      form: {},
      // 表单校验
      rules: {}
    }
  },
  created() {
    this.getList()
    this.getDicts('app_user_action_type').then(response => {
      this.actionTypeOptions = response.data
    })
    this.getDicts('app_user_by_type').then(response => {
      this.byTypeOptions = response.data
    })
  },
  methods: {
    /** 查询参数列表 */
    getList() {
      this.loading = true
      listUserOperLog(this.addDateRange(this.queryParams, this.dateRange)).then(response => {
        this.tableList = response.data.list
        this.total = response.data.count
        this.loading = false
      })
    },
    // 取消按钮
    cancel() {
      this.open = false
      this.reset()
    },
    // 表单重置
    reset() {
      this.form = {
      }
      this.resetForm('form')
    },

    actionTypeFormat(row) {
      return this.selectDictLabel(this.actionTypeOptions, row.actionType)
    },
    byTypeFormat(row) {
      return this.selectDictLabel(this.byTypeOptions, row.byType)
    },
    // 搜索按钮操作
    handleQuery() {
      this.queryParams.pageIndex = 1
      this.getList()
    },
    // 重置按钮操作
    resetQuery() {
      this.dateRange = []
      this.resetForm('queryForm')
      this.handleQuery()
    },
    /** 下载excel */
    handleExport() {
      this.$confirm('是否确认导出所选数据？', '提示', {
        confirmButtonText: '确认',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        exportUserOperLog(this.addDateRange(this.queryParams, this.dateRange)).then(response => {
          resolveBlob(response, '用户关键行为日志表')
        })
      }).catch(() => {})
    }
  }
}
</script>
