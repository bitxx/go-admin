<template>
  <BasicLayout>
    <template #wrapper>
      <el-card class="box-card">
        <el-form ref="queryForm" :model="queryParams" :inline="true" label-width="68px">
          <el-form-item label="用户编号" prop="id">
            <el-input v-model="queryParams.id" placeholder="请输入用户编号" clearable size="small" @keyup.enter.native="handleQuery" />
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
              v-permisaction="['admin:sysOperLog:export']"
              type="success"
              icon="el-icon-plus"
              size="mini"
              @click="handleExport"
            >Excel导出
            </el-button>
          </el-col>
        </el-row>

        <el-table v-loading="loading" stripe border :data="list">
          <el-table-column label="序号" type="index" align="center" width="80">
            <template slot-scope="scope">
              <span>{{ (queryParams.pageIndex - 1) * queryParams.pageSize + scope.$index + 1 }}</span>
            </template>
          </el-table-column>
          <el-table-column width="80" label="日志编号" align="center" prop="id" :show-overflow-tooltip="true" />
          <el-table-column width="100" label="用户编号" align="center" prop="userId" :show-overflow-tooltip="true" />
          <el-table-column width="100" label="请求方法" align="center" prop="requestMethod" :show-overflow-tooltip="true" />
          <el-table-column width="400" label="请求地址" align="center" prop="operUrl" :show-overflow-tooltip="true" />
          <el-table-column width="120" label="请求IP" align="center" prop="operIp" :show-overflow-tooltip="true" />
          <el-table-column width="200" label="访问位置" align="center" prop="operLocation" :show-overflow-tooltip="true" />
          <el-table-column width="120" label="返回码" align="center" prop="status" :show-overflow-tooltip="true" />
          <el-table-column width="200" label="返回数据" align="center" prop="jsonResult" :show-overflow-tooltip="true" />
          <el-table-column width="200" label="耗时" align="center" prop="latencyTime" :show-overflow-tooltip="true" />
          <el-table-column width="200" label="用户代理" align="center" prop="userAgent" :show-overflow-tooltip="true" />
          <el-table-column width="200" label="操作时间" align="center" prop="operTime" :show-overflow-tooltip="true">
            <template slot-scope="scope">
              <span>{{ parseTime(scope.row.operTime) }}</span>
            </template>
          </el-table-column>
          <el-table-column width="200" label="创建时间" align="center" prop="createdAt" :show-overflow-tooltip="true">
            <template slot-scope="scope">
              <span>{{ parseTime(scope.row.createdAt) }}</span>
            </template>
          </el-table-column>
          <el-table-column label="操作" width="150" fixed="right" align="center" class-name="small-padding fixed-width">
            <template slot-scope="scope">
              <el-button
                v-permisaction="['admin:sysOperLog:query']"
                size="mini"
                type="text"
                icon="el-icon-view"
                @click="handleView(scope.row,scope.index)"
              >详细</el-button>
              <el-button
                v-permisaction="['admin:sysOperLog:query']"
                size="mini"
                type="text"
                icon="el-icon-delete"
                @click="handleDelete(scope.row)"
              >删除</el-button>
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

        <!-- 操作日志详细 -->
        <el-dialog title="操作日志详细" :visible.sync="open" width="700px">
          <el-form ref="form" :model="form" label-width="100px" size="mini">
            <el-row>
              <el-col :span="24">
                <el-form-item label="请求地址：">{{ form.operUrl }}</el-form-item>
              </el-col>
              <el-col :span="12">
                <el-form-item
                  label="登录信息："
                >{{ form.operName }} / {{ form.operIp }} / {{ form.operLocation }}</el-form-item>
              </el-col>

              <el-col :span="12">
                <el-form-item label="请求方式：">{{ form.requestMethod }}</el-form-item>
              </el-col>
              <el-col :span="12">
                <el-form-item label="耗时：">{{ form.latencyTime }}</el-form-item>
              </el-col>
              <el-col :span="24">
                <el-form-item label="请求参数：">{{ form.operParam }}</el-form-item>
              </el-col>
              <el-col :span="24">
                <el-form-item label="返回参数：">{{ form.jsonResult }}</el-form-item>
              </el-col>
              <el-col :span="12">
                <el-form-item label="操作状态：">
                  <div v-if="form.status === '2'">正常</div>
                  <div v-else-if="form.status === '1'">关闭</div>
                </el-form-item>
              </el-col>
              <el-col :span="12">
                <el-form-item label="操作时间：">{{ parseTime(form.operTime) }}</el-form-item>
              </el-col>
              <el-col :span="24">
                <el-form-item v-if="form.status === 1" label="异常信息：">{{ form.errorMsg }}</el-form-item>
              </el-col>
            </el-row>
          </el-form>
          <div slot="footer" class="dialog-footer">
            <el-button @click="open = false">关 闭</el-button>
          </div>
        </el-dialog>
      </el-card>
    </template>
  </BasicLayout>
</template>

<script>
import { listSysOperlog, delSysOperlog, exportSysOperlog } from '@/api/sys/operlog'
import { resolveBlob } from '@/utils/download'

export default {
  name: 'SysOperalog',
  data() {
    return {
      // 遮罩层
      loading: true,
      // 总条数
      total: 0,
      // 表格数据
      list: [],
      // 是否显示弹出层
      open: false,
      // 日期范围
      dateRange: [],
      // 表单参数
      form: {},
      // 查询参数
      queryParams: {
        pageIndex: 1,
        pageSize: 10,
        title: undefined,
        operName: undefined,
        businessType: undefined
      }
    }
  },
  created() {
    this.getList()
  },
  methods: {
    /** 查询登录日志 */
    getList() {
      this.loading = true
      listSysOperlog(this.addDateRange(this.queryParams, this.dateRange)).then(response => {
        this.list = response.data.list
        this.total = response.data.count
        this.loading = false
      }
      )
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
    /** 详细按钮操作 */
    handleView(row) {
      this.open = true
      this.form = row
    },
    /** 删除按钮操作 */
    handleDelete(row) {
      const ids = [row.id]
      this.$confirm('是否确认删除日志编号为"' + ids + '"的数据项?', '警告', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(function() {
        return delSysOperlog({ 'ids': ids })
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
    /** 导出按钮操作 */
    /** 下载 */
    handleExport() {
      this.$confirm('是否确认导出所选数据？', '提示', {
        confirmButtonText: '确认',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        exportSysOperlog(this.addDateRange(this.queryParams, this.dateRange)).then(response => {
          resolveBlob(response, '登录日志')
        })
      }).catch(() => {
      })
    }
  }
}
</script>

