<template>
  <BasicLayout>
    <template #wrapper>
      <el-card class="box-card">
        <el-form ref="queryForm" :model="queryParams" :inline="true" label-width="68px">
          <el-form-item label="用户编号" prop="id">
            <el-input v-model="queryParams.id" placeholder="请输入用户编号" clearable size="small" @keyup.enter.native="handleQuery" />
          </el-form-item>
          <!--  <el-form-item label="等级编号" prop="levelId">
            <el-input v-model="queryParams.levelId" placeholder="请输入等级编号" clearable size="small" @keyup.enter.native="handleQuery" />
          </el-form-item>-->
          <el-form-item label="用户名" prop="userName">
            <el-input v-model="queryParams.userName" placeholder="请输入用户名" clearable size="small" @keyup.enter.native="handleQuery" />
          </el-form-item>
          <el-form-item label="真实姓名" prop="trueName">
            <el-input v-model="queryParams.trueName" placeholder="请输入真实姓名" clearable size="small" @keyup.enter.native="handleQuery" />
          </el-form-item>
          <el-form-item label="邮箱" prop="email">
            <el-input v-model="queryParams.email" placeholder="请输入电子邮箱" clearable size="small" @keyup.enter.native="handleQuery" />
          </el-form-item>
          <el-form-item label="手机号" prop="mobile">
            <el-input v-model="queryParams.mobile" placeholder="请输入手机号码" clearable size="small" @keyup.enter.native="handleQuery" />
          </el-form-item>
          <el-form-item label="上级用户邀请码" label-width="100" prop="parentRefCode">
            <el-input v-model="queryParams.parentRefCode" placeholder="请输入上级用户邀请码" clearable size="mini" @keyup.enter.native="handleQuery" />
          </el-form-item>
          <el-form-item label="当前用户邀请码" label-width="auto" prop="refCode">
            <el-input v-model="queryParams.refCode" placeholder="请输入当前用户邀请码" clearable size="mini" @keyup.enter.native="handleQuery" />
          </el-form-item>
          <el-form-item label="用户状态" prop="status">
            <el-select v-model="queryParams.status" placeholder="用户状态" clearable size="small">
              <el-option
                v-for="dict in statusOptions"
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
              v-permisaction="['app:user:user:add']"
              type="primary"
              icon="el-icon-plus"
              size="mini"
              @click="handleAdd"
            >
              新增
            </el-button>
          </el-col>
          <el-col :span="1.5">
            <el-button
              v-permisaction="['app:user:user:export']"
              type="success"
              icon="el-icon-download"
              size="mini"
              @click="handleExport"
            >
              Excel导出
            </el-button>
          </el-col>
        </el-row>

        <el-table v-loading="loading" stripe border show-summary :summary-method="getSummaries" :data="dataList">
          <el-table-column label="序号" type="index" align="center" width="80">
            <template slot-scope="scope">
              <span>{{ (queryParams.pageIndex - 1) * queryParams.pageSize + scope.$index + 1 }}</span>
            </template>
          </el-table-column>
          <el-table-column width="120" label="用户编号" align="center" prop="id" :show-overflow-tooltip="true" />
          <el-table-column width="120" label="等级编号" align="center" prop="levelId" :show-overflow-tooltip="true" />
          <el-table-column width="120" label="用户名" align="center" prop="userName" :show-overflow-tooltip="true" />
          <el-table-column width="120" label="真实姓名" align="center" prop="trueName" :show-overflow-tooltip="true" />
          <el-table-column width="120" label="账户余额" align="center" prop="money" :show-overflow-tooltip="true" />
          <el-table-column width="180" label="电子邮箱" align="center" prop="email" :show-overflow-tooltip="true" />
          <el-table-column label="国家区号" align="center" prop="mobileTitle" :show-overflow-tooltip="true" />
          <el-table-column width="160" label="手机号码" align="center" prop="mobile" :show-overflow-tooltip="true" />
          <el-table-column label="等级类型" align="center" prop="userLevel.levelType" :formatter="levelTypeFormat" width="100">
            <template slot-scope="scope">
              {{ levelTypeFormat(scope.row) }}
            </template>
          </el-table-column>
          <el-table-column width="60" label="等级" align="center" prop="userLevel.level" :show-overflow-tooltip="true" />
          <el-table-column width="120" label="当前用户邀请码" align="center" prop="refCode" :show-overflow-tooltip="true" />
          <el-table-column width="120" label="上级用户邀请码" align="center" prop="parentRefCode" :show-overflow-tooltip="true" />
          <el-table-column width="120" label="上级用户编号" align="center" prop="parentId" :show-overflow-tooltip="true" />
          <el-table-column label="用户状态" align="center" prop="status" :formatter="statusFormat" width="100">
            <template slot-scope="scope">
              {{ statusFormat(scope.row) }}
            </template>
          </el-table-column>
          <el-table-column width="200" label="创建时间" align="center" prop="createdAt" :show-overflow-tooltip="true">
            <template slot-scope="scope">
              <span>{{ parseTime(scope.row.createdAt) }}</span>
            </template>
          </el-table-column>
          <el-table-column width="60" fixed="right" label="操作" align="center" class-name="small-padding fixed-width">
            <template slot-scope="scope">
              <el-button
                v-permisaction="['app:user:user:edit']"
                size="mini"
                type="text"
                icon="el-icon-edit"
                @click="handleUpdate(scope.row)"
              >
                修改
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

        <!-- 添加或修改对话框 -->
        <el-dialog :close-on-click-modal="false" :title="title" :visible.sync="open" width="500" append-to-body>
          <el-form ref="form" :model="form" :rules="rules" label-width="120px">
            <el-row v-if="isEdit">
              <el-col :span="12">
                <el-form-item label="用户编号" prop="id">
                  <el-input v-model.number="form.id" placeholder="用户编号" :disabled="isEdit" />
                </el-form-item>
              </el-col>
              <el-col :span="12">
                <el-form-item label="等级编号" prop="levelId">
                  <el-input v-model.number="form.levelId" placeholder="等级编号" readonly @click.native="selectUserLevel">
                    <template slot="prepend">{{ levelName }}</template>
                  </el-input>
                </el-form-item>
              </el-col>
              <el-col :span="12">
                <el-form-item label="用户名" prop="userName">
                  <el-input v-model="form.userName" placeholder="用户名" />
                </el-form-item>
              </el-col>
              <el-col :span="12">
                <el-form-item label="真实姓名" prop="trueName">
                  <el-input v-model="form.trueName" placeholder="真实姓名" />
                </el-form-item>
              </el-col>
              <el-col :span="12">
                <el-form-item label="国家区号" prop="mobileTitle">
                  <el-input v-model="form.mobileTitle" placeholder="国家区号" />
                </el-form-item>
              </el-col>
              <el-col :span="12">
                <el-form-item label="手机号码" prop="mobile">
                  <el-input v-model="form.mobile" placeholder="手机号码" />
                </el-form-item>
              </el-col>
              <el-col :span="12">
                <el-form-item label="电子邮箱" prop="email">
                  <el-input v-model="form.email" placeholder="电子邮箱" />
                </el-form-item>
              </el-col>
            </el-row>
            <el-row v-if="!isEdit">
              <el-col :span="24">
                <el-form-item label="手机号区号" prop="mTitle">
                  <el-input v-model="form.mobileTitle" placeholder="请输入手机号区号" />
                </el-form-item>
              </el-col>
              <el-col :span="24">
                <el-form-item label="手机号码集合" prop="mobiles">
                  <el-input v-model="form.mobiles" type="textarea" rows="4" placeholder="请输入手机号(多个手机号之间使用半角逗号分割)" />
                </el-form-item>
              </el-col>
              <el-col :span="24">
                <el-form-item label="邮箱集合" prop="emails">
                  <el-input v-model="form.emails" type="textarea" rows="4" placeholder="请输入邮箱(多个邮箱之间使用半角逗号分割)" />
                </el-form-item>
              </el-col>
              <el-col :span="24">
                <el-form-item label="邀请码" prop="refCode">
                  <el-input v-model="form.refCode" placeholder="请输入邀请码" />
                </el-form-item>
              </el-col>
            </el-row>
          </el-form>
          <div slot="footer" class="dialog-footer">
            <el-button type="primary" @click="submitForm">确 定</el-button>
            <el-button @click="cancel">取 消</el-button>
          </div>
          <select-user-level ref="userLevel" @ok="getSelectUserLevelInfo" />
        </el-dialog>
      </el-card>
    </template>
  </BasicLayout>
</template>

<script>
import { addUser, getUser, listUser, updateUser, exportUser } from '@/api/app/user/user'
import { resolveBlob } from '@/utils/download'
import selectUserLevel from '@/views/app/select/selectUserLevel'
export default {
  name: 'User',
  components: { selectUserLevel },
  data() {
    return {
      levelName: '等级名称',
      // 遮罩层
      loading: true,
      // 总条数
      total: 0,
      // 弹出层标题
      title: '',
      // 是否显示弹出层
      open: false,
      isEdit: false,
      // 统计数据
      summaryData: {},
      // 日期范围
      dateRange: [],
      // 类型数据字典
      dataList: [],
      levelTypeOptions: [],
      statusOptions: [],
      // 查询参数
      queryParams: {
        pageIndex: 1,
        pageSize: 10,
        id: undefined,
        levelId: undefined,
        userName: undefined,
        trueName: undefined,
        email: undefined,
        mobile: undefined,
        parentId: undefined,
        refCode: undefined,
        parentRefCode: undefined
      },
      // 表单参数
      form: {},
      // 表单校验
      rules: {
        levelId: [{ required: true, message: '等级编号不能为空', trigger: 'blur' }],
        userName: [{ required: true, message: '用户名不能为空', trigger: 'blur' }],
        trueName: [{ required: true, message: '真实姓名不能为空', trigger: 'blur' }],
        email: [
          { required: true, message: '电子邮箱不能为空', trigger: 'blur' },
          { type: 'email', message: '请输入正确的邮箱地址', trigger: ['blur', 'change'] }
        ],
        mobileTitle: [{ required: true, message: '国家不能为空', trigger: 'blur' }],
        mobile: [
          { required: true, message: '手机号码不能为空', trigger: 'blur' },
          { pattern: /^[0-9]*[1-9][0-9]*$/, message: '请输入正确的手机号码', trigger: 'blur' }
        ]
      }
    }
  },
  created() {
    this.getList()
    this.getDicts('app_user_level_type').then(response => {
      this.levelTypeOptions = response.data
    })
    this.getDicts('sys_status').then(response => {
      this.statusOptions = response.data
    })
  },
  methods: {
    /** 查询参数列表 */
    getList() {
      this.loading = true
      listUser(this.addDateRange(this.queryParams, this.dateRange)).then(response => {
        this.dataList = response.data.list
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
        id: undefined,
        levelId: undefined,
        userName: undefined,
        trueName: undefined,
        email: undefined,
        mobileTitle: undefined,
        mobile: undefined,
        mobiles: undefined,
        emails: undefined,
        refCode: undefined
      }
      this.resetForm('form')
    },
    levelTypeFormat(row) {
      if (row.userLevel === null) {
        return ''
      }
      return this.selectDictLabel(this.levelTypeOptions, row.userLevel.levelType)
    },
    statusFormat(row) {
      return this.selectDictLabel(this.statusOptions, row.status)
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
    // 新增按钮操作
    handleAdd() {
      this.reset()
      this.open = true
      this.title = '添加用户管理'
      this.isEdit = false
    },
    // 修改按钮操作
    handleUpdate(row) {
      this.reset()
      getUser(row.id).then(response => {
        this.form = response.data
        this.open = true
        this.title = '修改用户管理'
        this.isEdit = true
      })
    },
    /** 提交按钮 */
    submitForm: function() {
      this.$refs['form'].validate(valid => {
        if (valid) {
          if (this.form.id !== undefined) {
            updateUser(this.form).then(response => {
              if (response.code === 200) {
                this.msgSuccess(response.msg)
                this.open = false
                this.getList()
              } else {
                this.msgError(response.msg)
              }
            })
          } else {
            addUser(this.form).then(response => {
              if (response.code === 200) {
                this.msgSuccess(response.msg)
                this.open = false
                this.getList()
              } else {
                this.msgError(response.msg)
              }
            })
          }
        }
      })
    },
    /** 下载excel */
    handleExport() {
      this.$confirm('是否确认导出所选数据？', '提示', {
        confirmButtonText: '确认',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        exportUser(this.addDateRange(this.queryParams, this.dateRange)).then(response => {
          resolveBlob(response, '用户管理')
        })
      }).catch(() => {})
    },
    /** 显示选项框selectUserLevel */
    selectUserLevel() {
      this.$refs.userLevel.show('1')
    },
    /** 获取返回的信息 */
    getSelectUserLevelInfo(row) {
      this.form.levelId = row.id
      this.levelName = row.name.toString()
      this.$refs['form'].validateField('levelId')
    },
    /** 统计 */
    getSummaries(param) {
      const { columns } = param
      const sums = []
      columns.forEach((column, index) => {
        if (index === 1) {
          sums[index] = '合计'
          return
        }
        if (column.property === 'money') {
          sums[index] = this.summaryData.money
        }
      })
      return sums
    }
  }
}
</script>
