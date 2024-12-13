<template>
  <BasicLayout>
    <template #wrapper>
      <el-card class="box-card">
        <el-form ref="queryForm" :model="queryParams" :inline="true" label-width="68px">
          <el-form-item label-width="100" label="版本号" prop="version">
            <el-input v-model="queryParams.version" placeholder="请输入版本号" clearable size="small" @keyup.enter.native="handleQuery" />
          </el-form-item>
          <el-form-item label-width="100" label="平台" prop="platform">
            <el-select v-model="queryParams.platform" placeholder="平台" clearable size="small">
              <el-option
                v-for="dict in platformOptions"
                :key="dict.dictValue"
                :label="dict.dictLabel"
                :value="dict.dictValue"
              />
            </el-select>
          </el-form-item>
          <el-form-item label-width="100" label="App类型" prop="appType">
            <el-select v-model="queryParams.appType" placeholder="App类型" clearable size="small">
              <el-option
                v-for="dict in appTypeOptions"
                :key="dict.dictValue"
                :label="dict.dictLabel"
                :value="dict.dictValue"
              />
            </el-select>
          </el-form-item>
          <el-form-item label-width="100" label="下载方式" prop="downloadType">
            <el-select v-model="queryParams.downloadType" placeholder="下载方式" clearable size="small">
              <el-option
                v-for="dict in downloadTypeOptions"
                :key="dict.dictValue"
                :label="dict.dictLabel"
                :value="dict.dictValue"
              />
            </el-select>
          </el-form-item>
          <el-form-item label-width="100" label="发布状态" prop="status">
            <el-select v-model="queryParams.status" placeholder="发布状态" clearable size="small">
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
              v-permisaction="['plugins:filemgr-app:add']"
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
              v-permisaction="['plugins:filemgr-app:export']"
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
          <el-table-column width="100" label="App编号" align="center" prop="id" :show-overflow-tooltip="true" />
          <el-table-column width="100" label="版本号" align="center" prop="version" :show-overflow-tooltip="true" />
          <el-table-column width="100" label="系统平台" align="center" prop="platform" :formatter="platformFormat">
            <template slot-scope="scope">
              {{ platformFormat(scope.row) }}
            </template>
          </el-table-column>
          <el-table-column width="150" label="版本类型" align="center" prop="appType" :formatter="appTypeFormat">
            <template slot-scope="scope">
              {{ appTypeFormat(scope.row) }}
            </template>
          </el-table-column>
          <el-table-column width="100" label="下载类型" align="center" prop="downloadType" :formatter="downloadTypeFormat">
            <template slot-scope="scope">
              {{ downloadTypeFormat(scope.row) }}
            </template>
          </el-table-column>
          <el-table-column width="100" label="发布状态" align="center" prop="status" :formatter="statusFormat">
            <template slot-scope="scope">
              {{ statusFormat(scope.row) }}
            </template>
          </el-table-column>
          <el-table-column width="300" label="下载地址" align="center" prop="downloadUrl" :show-overflow-tooltip="true" />
          <el-table-column width="300" label="服务器本地地址" align="center" prop="localAddress" :show-overflow-tooltip="true" />
          <el-table-column width="500" label="更新内容" align="center" prop="remark" :show-overflow-tooltip="true" />
          <el-table-column width="200" label="创建时间" align="center" prop="createdAt" :show-overflow-tooltip="true">
            <template slot-scope="scope">
              <span>{{ parseTime(scope.row.createdAt) }}</span>
            </template>
          </el-table-column>
          <el-table-column width="120" fixed="right" label="操作" align="center" class-name="small-padding fixed-width">
            <template slot-scope="scope">
              <el-button
                v-permisaction="['plugins:filemgr-app:edit']"
                size="mini"
                type="text"
                icon="el-icon-edit"
                @click="handleUpdate(scope.row)"
              >
                修改
              </el-button>
              <el-button
                v-permisaction="['plugins:filemgr-app:del']"
                size="mini"
                type="text"
                icon="el-icon-delete"
                @click="handleDelete(scope.row)"
              >
                删除
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
          <el-form ref="form" :model="form" :rules="rules" label-width="80px">
            <el-row v-if="!isEdit">
              <el-col :span="12">
                <el-form-item label="版本号" prop="version">
                  <el-input v-model="form.version" placeholder="版本号" />
                </el-form-item>
              </el-col>
              <el-col :span="12">
                <el-form-item label="平台" prop="platform">
                  <el-select v-model="form.platform" placeholder="请选择">
                    <el-option
                      v-for="dict in platformOptions"
                      :key="dict.dictValue"
                      :label="dict.dictLabel"
                      :value="dict.dictValue"
                    />
                  </el-select>
                </el-form-item>
              </el-col>
              <el-col :span="12">
                <el-form-item label="下载类型" prop="downloadType">
                  <el-select v-model="form.downloadType" placeholder="请选择">
                    <el-option
                      v-for="dict in downloadTypeOptions"
                      :key="dict.dictValue"
                      :label="dict.dictLabel"
                      :value="dict.dictValue"
                    />
                  </el-select>
                </el-form-item>
              </el-col>
              <el-col :span="12">
                <el-form-item label="App类型" prop="appType">
                  <el-select v-model="form.appType" placeholder="请选择">
                    <el-option
                      v-for="dict in appTypeOptions"
                      :key="dict.dictValue"
                      :label="dict.dictLabel"
                      :value="dict.dictValue"
                    />
                  </el-select>
                </el-form-item>
              </el-col>
              <el-col v-if="form.downloadType==='2'" :span="24">
                <el-form-item label="下载地址" prop="downloadUrl">
                  <el-input v-model="form.downloadUrl" placeholder="下载地址" />
                </el-form-item>
              </el-col>
              <el-col v-if="form.downloadType==='1'" :span="24">
                <el-form-item label="本地根URL" prop="localRootUrl">
                  <el-input v-model="form.localRootUrl" placeholder="本地根URL" />
                </el-form-item>
              </el-col>
              <el-col v-if="form.downloadType==='1' || form.downloadType==='3'" :span="24">
                <el-form-item label="上传App">
                  <el-upload
                    ref="upload"
                    :limit="1"
                    accept=".apk, .ipa"
                    :headers="upload.headers"
                    :action="upload.url"
                    :disabled="upload.isUploading"
                    :on-progress="handleFileUploadProgress"
                    :on-success="handleFileSuccess"
                    :show-file-list="true"
                    :auto-upload="true"
                    drag
                  >
                    <i class="el-icon-upload" />
                    <div class="el-upload__text">
                      将文件拖到此处，或<em>点击上传</em>
                    </div>
                    <div slot="tip" class="el-upload__tip" style="color:red">提示：仅允许导入“apk”或“ipa”格式文件！</div>
                  </el-upload>
                </el-form-item>
              </el-col>
              <el-col :span="24">
                <el-form-item label="更新内容" prop="remark">
                  <el-input
                    v-model="form.remark"
                    type="textarea"
                    :rows="4"
                    placeholder="请输入更新内容"
                  />
                </el-form-item>
              </el-col>
            </el-row>
            <el-row v-if="isEdit">
              <el-col :span="12">
                <el-form-item label="App状态" prop="status">
                  <el-select v-model="form.status" placeholder="请选择">
                    <el-option
                      v-for="dict in statusOptions"
                      :key="dict.dictValue"
                      :label="dict.dictLabel"
                      :value="dict.dictValue"
                    />
                  </el-select>
                </el-form-item>
              </el-col>
            </el-row>
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
import { getToken } from '@/utils/auth'
import { addFilemgrApp, delFilemgrApp, getFilemgrApp, listFilemgrApp, updateFilemgrApp, exportFilemgrApp } from '@/api/plugins/filemgr/filemgr-app'
import { resolveBlob } from '@/utils/download'
export default {
  name: 'FilemgrApp',
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
      platformOptions: [],
      appTypeOptions: [],
      downloadTypeOptions: [],
      statusOptions: [],
      // 查询参数
      queryParams: {
        pageIndex: 1,
        pageSize: 10,
        version: undefined,
        platform: undefined,
        appType: undefined,
        status: undefined,
        downloadType: undefined,
        downloadUrl: undefined
      },
      // 用户导入参数
      upload: {
        // 是否禁用上传
        isUploading: false,
        // 设置上传的请求头部
        headers: { 'Authorization': 'Bearer ' + getToken() },
        // 上传的地址
        url: process.env.VUE_APP_BASE_API + '/admin-api/v1/plugins/filemgr/filemgr-app/upload'
      },
      // 表单参数
      form: {},
      // 表单校验
      rules: {
        filePath: [{ required: true, message: '文件路径不能为空', trigger: 'blur' }],
        version: [{ required: true, message: '版本号不能为空', trigger: 'blur' }],
        platform: [{ required: true, message: '平台不能为空', trigger: 'blur' }],
        appType: [{ required: true, message: 'App类型不能为空', trigger: 'blur' }],
        downloadType: [{ required: true, message: '下载类型不能为空', trigger: 'blur' }],
        localRootUrl: [{ required: true, message: '本地根URL不得为空', trigger: 'blur' }],
        remark: [{ required: true, message: '更新内容不能为空', trigger: 'blur' }]
      }
    }
  },
  created() {
    this.getList()
    this.getDicts('app_platform').then(response => {
      this.platformOptions = response.data
    })
    this.getDicts('plugin_filemgr_app_type').then(response => {
      this.appTypeOptions = response.data
    })
    this.getDicts('plugin_filemgr_app_download_type').then(response => {
      this.downloadTypeOptions = response.data
    })
    this.getDicts('plugin_filemgr_app_publish_status').then(response => {
      this.statusOptions = response.data
    })
  },
  methods: {
    /** 查询参数列表 */
    getList() {
      this.loading = true
      listFilemgrApp(this.addDateRange(this.queryParams, this.dateRange))
        .then(response => {
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
        id: undefined,
        version: undefined,
        platform: undefined,
        appType: undefined,
        remark: undefined,
        downloadType: undefined,
        downloadUrl: undefined,
        localRootUrl: undefined,
        localAddress: undefined
      }
      this.resetForm('form')
    },
    platformFormat(row) {
      return this.selectDictLabel(this.platformOptions, row.platform)
    },
    appTypeFormat(row) {
      return this.selectDictLabel(this.appTypeOptions, row.appType)
    },
    downloadTypeFormat(row) {
      return this.selectDictLabel(this.downloadTypeOptions, row.downloadType)
    },
    statusFormat(row) {
      return this.selectDictLabel(this.statusOptions, row.status)
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
    /** 新增按钮操作 */
    handleAdd() {
      this.reset()
      this.open = true
      this.title = '新增App版本'
      this.isEdit = false
    },
    // 修改按钮操作
    handleUpdate(row) {
      this.reset()
      getFilemgrApp(row.id).then(response => {
        this.form = response.data
        this.open = true
        this.title = '修改App管理'
        this.isEdit = true
      })
    },
    /** 提交按钮 */
    submitForm: function() {
      this.$refs['form'].validate(valid => {
        if (valid) {
          if (this.form.id !== undefined) {
            updateFilemgrApp(this.form).then(response => {
              if (response.code === 200) {
                this.msgSuccess(response.msg)
                this.open = false
                this.getList()
              } else {
                this.msgError(response.msg)
              }
            })
          } else {
            addFilemgrApp(this.form).then(response => {
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
    // 删除按钮操作
    handleDelete(row) {
      const ids = [row.id]
      this.$confirm('是否确认删除编号为"' + ids + '"的数据项?', '警告', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(function() {
        return delFilemgrApp({ 'ids': ids })
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
    /** 下载excel */
    handleExport() {
      this.$confirm('是否确认导出所选数据？', '提示', {
        confirmButtonText: '确认',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        exportFilemgrApp(this.addDateRange(this.queryParams, this.dateRange)).then(response => {
          resolveBlob(response, 'App管理')
        })
      }).catch(() => {})
    },
    // 文件上传中处理
    handleFileUploadProgress(event, file, fileList) {
      this.upload.isUploading = true
    },
    // 文件上传成功处理
    handleFileSuccess(response, file, fileList) {
      this.upload.isUploading = false
      if (response.code !== 200 || response.data === undefined || response.data === '') {
        return
      }
      this.form.localAddress = response.data
    }
  }
}
</script>
