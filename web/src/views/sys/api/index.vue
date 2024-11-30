<template>
  <BasicLayout>
    <template #wrapper>
      <el-card class="box-card">
        <el-form ref="queryForm" :model="queryParams" :inline="true" label-width="68px">

          <el-form-item label-width="100" label="标题" prop="title">
            <el-input v-model="queryParams.title" placeholder="请输入标题" clearable size="small" @keyup.enter.native="handleQuery" />
          </el-form-item>

          <el-form-item label-width="100" label="地址" prop="path">
            <el-input v-model="queryParams.path" placeholder="请输入地址" clearable size="small" @keyup.enter.native="handleQuery" />
          </el-form-item>

          <el-form-item label-width="100" label="请求方法" prop="action">
            <el-select v-model="queryParams.action" placeholder="请求方法" clearable size="small">
              <el-option
                v-for="dict in actionOptions"
                :key="dict.dictValue"
                :label="dict.dictLabel"
                :value="dict.dictValue"
              />
            </el-select>
          </el-form-item>

          <el-form-item label-width="100" label="接口类型" prop="apiType">
            <el-select v-model="queryParams.apiType" placeholder="接口类型" clearable size="small">
              <el-option
                v-for="dict in apiTypeOptions"
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
              v-permisaction="['sys:api:export']"
              type="success"
              icon="el-icon-plus"
              size="mini"
              @click="handleExport"
            >Excel导出
            </el-button>
          </el-col>
        </el-row>

        <el-table v-loading="loading" stripe border :data="sysapiList">
          <el-table-column label="序号" type="index" align="center" width="80">
            <template slot-scope="scope">
              <span>{{ (queryParams.pageIndex - 1) * queryParams.pageSize + scope.$index + 1 }}</span>
            </template>
          </el-table-column>
          <el-table-column width="120" label="接口编号" align="center" prop="id" :show-overflow-tooltip="true" />
          <el-table-column width="240" label="标题" align="center" prop="title" :show-overflow-tooltip="true" />
          <el-table-column width="100" label="接口类型" align="center" prop="type" :formatter="apiTypeFormat">
            <template slot-scope="scope">
              {{ apiTypeFormat(scope.row) }}
            </template>
          </el-table-column>
          <el-table-column width="300" label="Handle" align="center" prop="handle" :show-overflow-tooltip="true" />
          <el-table-column width="300" label="请求地址" align="center" prop="path" :show-overflow-tooltip="true" />
          <el-table-column width="100" label="请求方法" align="center" prop="action" :formatter="actionFormat">
            <template slot-scope="scope">
              {{ actionFormat(scope.row) }}
            </template>
          </el-table-column>
          <el-table-column width="200" label="更新时间" align="center" prop="updatedAt" :show-overflow-tooltip="true">
            <template slot-scope="scope">
              <span>{{ parseTime(scope.row.updatedAt) }}</span>
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
                v-permisaction="['content:announcement:edit']"
                size="mini"
                type="text"
                icon="el-icon-edit"
                @click="handleUpdate(scope.row)"
              >修改
              </el-button>
              <el-button
                v-permisaction="['content:announcement:remove']"
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

        <!-- 添加或修改对话框 -->
        <el-dialog :close-on-click-modal="false" :title="title" :visible.sync="open" width="2200" append-to-body>
          <el-form ref="form" :model="form" :rules="rules" label-width="80px">
            <el-form-item label="Handle" prop="handle">
              <el-input v-model="form.handle" placeholder="handle" :disabled="isEdit" />
            </el-form-item>
            <el-form-item label="请求地址" prop="path">
              <el-input v-model="form.path" placeholder="path" :disabled="isEdit" />
            </el-form-item>
            <el-form-item label="请求方法" prop="action">
              <el-select v-model="form.action" placeholder="请选择" :disabled="isEdit">
                <el-option
                  v-for="dict in actionOptions"
                  :key="dict.dictValue"
                  :label="dict.dictLabel"
                  :value="dict.dictValue"
                />
              </el-select>
            </el-form-item>
            <el-form-item label="标题" prop="title">
              <el-input v-model="form.title" placeholder="标题" />
            </el-form-item>
            <el-form-item label="接口类型" prop="apiType">
              <el-select v-model="form.apiType" placeholder="请选择">
                <el-option
                  v-for="dict in apiTypeOptions"
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
import { delSysApi, getSysApi, listSysApi, updateSysApi, exportSysApi } from '@/api/sys/api'
import { resolveBlob } from '@/utils/download'

export default {
  name: 'SysApi',
  components: {
  },
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
      // 是否编辑
      isEdit: false,
      // 日期范围
      dateRange: [],
      // 类型数据字典
      apiTypeOptions: [],
      actionOptions: [],
      sysapiList: [],

      // 查询参数
      queryParams: {
        pageIndex: 1,
        pageSize: 10,
        title: undefined,
        path: undefined,
        action: undefined
      },
      // 表单参数
      form: {},
      // 表单校验
      rules: {
        title: [{ required: true, message: '标题不能为空', trigger: 'blur' }],
        // action: [{ required: true, message: '类型不能为空', trigger: 'blur' }],
        apiType: [{ required: true, message: '接口类型不得为空', trigger: 'blur' }]
      }
    }
  },
  created() {
    this.getList()
    this.getDicts('sys_api_type').then(response => {
      this.apiTypeOptions = response.data
    })
    this.getDicts('sys_api_action').then(response => {
      this.actionOptions = response.data
    })
  },
  methods: {
    /** 查询参数列表 */
    getList() {
      this.loading = true
      listSysApi(this.addDateRange(this.queryParams, this.dateRange)).then(response => {
        this.sysapiList = response.data.list
        this.total = response.data.count
        this.loading = false
      }
      )
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
        title: undefined,
        path: undefined,
        paths: undefined,
        action: undefined
      }
      this.resetForm('form')
    },
    apiTypeFormat(row) {
      return this.selectDictLabel(this.apiTypeOptions, row.apiType)
    },
    actionFormat(row) {
      return this.selectDictLabel(this.actionOptions, row.action)
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
    },
    /** 修改按钮操作 */
    handleUpdate(row) {
      this.reset()
      const id = row.id
      getSysApi(id).then(response => {
        this.form = response.data
        this.open = true
        this.title = '修改接口管理'
        this.isEdit = true
      })
    },
    /** 提交按钮 */
    submitForm: function() {
      this.$refs['form'].validate(valid => {
        if (valid) {
          if (this.form.id !== undefined) {
            updateSysApi(this.form).then(response => {
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
    /** 删除按钮操作 */
    handleDelete(row) {
      const ids = [row.id]
      this.$confirm('是否确认删除编号为"' + ids + '"的数据项?', '警告', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(function() {
        return delSysApi({ 'ids': ids })
      }).then((response) => {
        if (response.code === 200) {
          this.msgSuccess(response.msg)
          this.open = false
          this.getList()
        } else {
          this.msgError(response.msg)
        }
      }).catch(function() {})
    },
    /** 下载 */
    handleExport() {
      this.$confirm('是否确认导出所选数据？', '提示', {
        confirmButtonText: '确认',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        exportSysApi(this.addDateRange(this.queryParams, this.dateRange)).then(response => {
          resolveBlob(response, 'api列表')
        })
      }).catch(() => {
      })
    }
  }
}
</script>
