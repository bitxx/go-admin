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
          <el-form-item label-width="200" label="是否允许登录" prop="canLogin">
            <el-select v-model="queryParams.canLogin" placeholder="是否允许登录" clearable size="small">
              <el-option
                v-for="dict in canLoginOptions"
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

        <el-table v-loading="loading" stripe border :data="tableList">
          <el-table-column label="序号" type="index" align="center" width="80">
            <template slot-scope="scope">
              <span>{{ (queryParams.pageIndex - 1) * queryParams.pageSize + scope.$index + 1 }}</span>
            </template>
          </el-table-column>
          <el-table-column width="100" label="配置编号" align="center" prop="id" :show-overflow-tooltip="true" />
          <el-table-column width="100" label="用户编号" align="center" prop="userId" :show-overflow-tooltip="true" />
          <el-table-column width="100" label="用户名称" align="center" prop="user.userName" :show-overflow-tooltip="true" />
          <el-table-column width="150" label="用户邮箱" align="center" prop="user.email" :show-overflow-tooltip="true" />
          <el-table-column width="150" label="用户手机号" align="center" prop="user.mobile" :show-overflow-tooltip="true" />
          <el-table-column label="是否允许登录" align="center" prop="canLogin" :formatter="canLoginFormat">
            <template slot-scope="scope">
              {{ canLoginFormat(scope.row) }}
            </template>
          </el-table-column>
          <el-table-column width="200" label="创建时间" align="center" prop="createdAt" :show-overflow-tooltip="true">
            <template slot-scope="scope">
              <span>{{ parseTime(scope.row.createdAt) }}</span>
            </template>
          </el-table-column>
          <el-table-column width="160" fixed="right" label="操作" align="center" class-name="small-padding fixed-width">
            <template slot-scope="scope">
              <el-button
                v-permisaction="['app:user:user-conf:edit']"
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
        <el-dialog :close-on-click-modal="false" :title="title" :visible.sync="open" width="300" append-to-body>
          <el-form ref="form" :model="form" :rules="rules" label-width="120px">
            <el-form-item label="用户编号" prop="userId">
              <el-input v-model.number="form.userId" placeholder="用户编号" :disabled="isEdit" />
            </el-form-item>
            <el-form-item label="是否允许登录" prop="canLogin">
              <el-select v-model="form.canLogin" placeholder="请选择">
                <el-option
                  v-for="dict in canLoginOptions"
                  :key="dict.dictValue"
                  :label="dict.dictLabel"
                  :value="dict.dictValue"
                />
              </el-select>
            </el-form-item>
          </el-form>
          <div slot="footer" class="dialog-footer">
            <el-button type="primary" @click="submitForm">确 定</el-button>
            <el-button @click="cancel">取 消</el-button>
          </div>
        </el-dialog>
      </el-card>
    </template>
  </BasicLayout>
</template>

<script>
import { getUserConf, listUserConf, updateUserConf } from '@/api/app/user/user-conf'
export default {
  name: 'UserConf',
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
      canLoginOptions: [],
      // 查询参数
      queryParams: {
        pageIndex: 1,
        pageSize: 10,
        userId: undefined,
        canLogin: undefined
      },
      // 表单参数
      form: {},
      // 表单校验
      rules: {
        userId: [{ required: true, message: '用户编号不能为空', trigger: 'blur' }],
        canLogin: [{ required: true, message: '是否允许登录不能为空', trigger: 'blur' }]
      }
    }
  },
  created() {
    this.getList()
    this.getDicts('sys_yes_no').then(response => {
      this.canLoginOptions = response.data
    })
  },
  methods: {
    /** 查询参数列表 */
    getList() {
      this.loading = true
      listUserConf(this.addDateRange(this.queryParams, this.dateRange)).then(response => {
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
        userId: undefined,
        canLogin: undefined
      }
      this.resetForm('form')
    },

    canLoginFormat(row) {
      return this.selectDictLabel(this.canLoginOptions, row.canLogin)
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
    // 修改按钮操作
    handleUpdate(row) {
      this.reset()
      getUserConf(row.id).then(response => {
        this.form = response.data
        this.open = true
        this.title = '修改用户配置'
        this.isEdit = true
      })
    },
    /** 提交按钮 */
    submitForm: function() {
      this.$refs['form'].validate(valid => {
        if (valid) {
          if (this.form.id !== undefined) {
            updateUserConf(this.form).then(response => {
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
    }
  }
}
</script>
