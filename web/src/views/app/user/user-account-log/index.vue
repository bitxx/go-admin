<template>
  <BasicLayout>
    <template #wrapper>
      <el-card class="box-card">
        <el-form ref="queryForm" :model="queryParams" :inline="true" label-width="68px">
          <el-form-item label="用户编号" prop="userId">
            <el-input v-model="queryParams.userId" placeholder="请输入用户编号" clearable size="small" @keyup.enter.native="handleQuery" />
          </el-form-item>
          <el-form-item label="用户名称" label-width="100" prop="userName">
            <el-input v-model="queryParams.userName" placeholder="请输入用户名称" clearable size="mini" @keyup.enter.native="handleQuery" />
          </el-form-item>
          <el-form-item label-width="100" label="用户邮箱" prop="email">
            <el-input v-model="queryParams.email" placeholder="请输入用户邮箱" clearable size="small" @keyup.enter.native="handleQuery" />
          </el-form-item>
          <el-form-item label-width="100" label="用户手机号" prop="mobile">
            <el-input v-model="queryParams.mobile" placeholder="请输入用户手机号" clearable size="small" @keyup.enter.native="handleQuery" />
          </el-form-item>
          <el-form-item label="金额类型" prop="moneyType">
            <el-select v-model="queryParams.moneyType" placeholder="账变记录金额类型" clearable size="small">
              <el-option
                v-for="dict in moneyTypeOptions"
                :key="dict.dictValue"
                :label="dict.dictLabel"
                :value="dict.dictValue"
              />
            </el-select>
          </el-form-item>
          <el-form-item label="帐变类型" prop="changeType">
            <el-select v-model="queryParams.changeType" placeholder="账变记录帐变类型" clearable size="small">
              <el-option
                v-for="dict in changeTypeOptions"
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
              v-permisaction="['app:user:user-account-log:export']"
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
          <el-table-column width="100" label="账变编号" align="center" prop="id" :show-overflow-tooltip="true" />
          <el-table-column width="100" label="用户编号" align="center" prop="userId" :show-overflow-tooltip="true" />
          <el-table-column width="100" label="用户名称" align="center" prop="user.userName" :show-overflow-tooltip="true" />
          <el-table-column width="150" label="用户邮箱" align="center" prop="user.email" :show-overflow-tooltip="true" />
          <el-table-column width="150" label="用户手机号" align="center" prop="user.mobile" :show-overflow-tooltip="true" />
          <el-table-column label="账变金额" align="center" prop="changeMoney" :show-overflow-tooltip="true" />
          <el-table-column label="账变前金额" align="center" prop="beforeMoney" :show-overflow-tooltip="true" />
          <el-table-column label="账变后金额" align="center" prop="afterMoney" :show-overflow-tooltip="true" />
          <el-table-column label="金额类型" align="center" prop="moneyType" :formatter="moneyTypeFormat" width="100">
            <template slot-scope="scope">
              {{ moneyTypeFormat(scope.row) }}
            </template>
          </el-table-column>
          <el-table-column label="帐变类型" align="center" prop="changeType" :formatter="changeTypeFormat" width="100">
            <template slot-scope="scope">
              {{ changeTypeFormat(scope.row) }}
            </template>
          </el-table-column>
          <el-table-column width="200" label="创建时间" align="center" prop="createdAt" :show-overflow-tooltip="true">
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
      </el-card>
    </template>
  </BasicLayout>
</template>

<script>
import { listUserAccountLog, exportUserAccountLog } from '@/api/app/user/user-account-log'
import { resolveBlob } from '@/utils/download'
export default {
  name: 'UserAccountLog',
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
      moneyTypeOptions: [],
      changeTypeOptions: [],
      // 查询参数
      queryParams: {
        pageIndex: 1,
        pageSize: 10,
        userId: undefined,
        moneyType: undefined,
        changeType: undefined
      },
      // 表单参数
      form: {},
      // 表单校验
      rules: {}
    }
  },
  created() {
    this.getList()
    this.getDicts('app_money_type').then(response => {
      this.moneyTypeOptions = response.data
    })
    this.getDicts('app_account_change_type').then(response => {
      this.changeTypeOptions = response.data
    })
  },
  methods: {
    /** 查询参数列表 */
    getList() {
      this.loading = true
      listUserAccountLog(this.addDateRange(this.queryParams, this.dateRange)).then(response => {
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

    moneyTypeFormat(row) {
      return this.selectDictLabel(this.moneyTypeOptions, row.moneyType)
    },
    changeTypeFormat(row) {
      return this.selectDictLabel(this.changeTypeOptions, row.changeType)
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
        exportUserAccountLog(this.addDateRange(this.queryParams, this.dateRange)).then(response => {
          resolveBlob(response, '账变记录')
        })
      }).catch(() => {})
    }
  }
}
</script>
