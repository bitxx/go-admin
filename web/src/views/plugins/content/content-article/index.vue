<template>
  <BasicLayout>
    <template #wrapper>
      <el-card class="box-card">
        <el-form ref="queryForm" :model="queryParams" :inline="true" label-width="68px">
          <el-form-item label-width="100" label="文章编号" prop="id">
            <el-input v-model="queryParams.id" placeholder="请输入文章编号" clearable size="small" @keyup.enter.native="handleQuery" />
          </el-form-item>
          <el-form-item label-width="100" label="分类编号" prop="cateId">
            <el-input v-model="queryParams.cateId" placeholder="请输入分类编号" clearable size="small" @keyup.enter.native="handleQuery" />
          </el-form-item>
          <el-form-item label-width="100" label="标题" prop="name">
            <el-input v-model="queryParams.name" placeholder="请输入标题" clearable size="small" @keyup.enter.native="handleQuery" />
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
              v-permisaction="['plugins:content:content-article:add']"
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
              v-permisaction="['plugins:content:content-article:export']"
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
          <el-table-column width="100" label="文章编号" align="center" prop="id" :show-overflow-tooltip="true" />
          <el-table-column width="100" label="分类编号" align="center" prop="cateId" :show-overflow-tooltip="true" />
          <el-table-column width="100" label="分类名称" align="center" prop="contentCategory.name" :show-overflow-tooltip="true" />
          <el-table-column label="标题" align="center" prop="name" :show-overflow-tooltip="true" />
          <el-table-column width="200" label="文章内容" align="center" prop="content" :show-overflow-tooltip="true" />
          <el-table-column label="备注信息" align="center" prop="remark" :show-overflow-tooltip="true" />
          <el-table-column label="状态" align="center" prop="status" :formatter="statusFormat" width="100">
            <template slot-scope="scope">
              {{ statusFormat(scope.row) }}
            </template>
          </el-table-column>
          <el-table-column width="100" label="更新人编号" align="center" prop="updateBy" :show-overflow-tooltip="true" />
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
          <el-table-column width="140" fixed="right" label="操作" align="center" class-name="small-padding fixed-width">
            <template slot-scope="scope">
              <el-button
                v-permisaction="['plugins:content:content-article:edit']"
                size="mini"
                type="text"
                icon="el-icon-edit"
                @click="handleUpdate(scope.row)"
              >
                修改
              </el-button>
              <el-button
                v-permisaction="['plugins:content:content-article:del']"
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
            <el-form-item label="分类编号" prop="cateId">
              <el-input v-model.number="form.cateId" placeholder="分类编号" readonly @click.native="selectCategory">
                <template slot="prepend">{{ categoryName }}</template>
              </el-input>
            </el-form-item>
            <el-form-item label="标题" prop="name">
              <el-input v-model="form.name" placeholder="标题" />
            </el-form-item>
            <el-form-item label="状态" prop="status">
              <el-select v-model="form.status" placeholder="请选择">
                <el-option
                  v-for="dict in statusOptions"
                  :key="dict.dictValue"
                  :label="dict.dictLabel"
                  :value="dict.dictValue"
                />
              </el-select>
            </el-form-item>
            <el-form-item v-if="open" label="文本内容" prop="content">
              <Tinymce ref="editor" v-model="form.content" :height="200" />
            </el-form-item>
            <el-form-item label="备注信息" prop="remark">
              <el-input v-model="form.remark" placeholder="备注信息" />
            </el-form-item>
          </el-form>
          <div slot="footer" class="dialog-footer">
            <el-button type="primary" @click="submitForm">确 定</el-button>
            <el-button @click="cancel">取 消</el-button>
          </div>
          <select-category ref="category" @ok="getCategorySelectInfo" />
        </el-dialog>
      </el-card>
    </template>
  </BasicLayout>
</template>

<script>
import { addContentArticle, delContentArticle, getContentArticle, listContentArticle, updateContentArticle, exportContentArticle } from '@/api/plugins/content/content-article'
import Tinymce from '@/components/Tinymce/index.vue'
import { resolveBlob } from '@/utils/download'
import selectCategory from '@/views/plugins/content/select/selectCategory.vue'

const defaultForm = {
  status: 'draft'
}
export default {
  name: 'ContentArticle',
  components: { Tinymce, selectCategory },
  data() {
    return {
      categoryName: '请选择',
      // 遮罩层
      loading: true,
      postForm: Object.assign({}, defaultForm),
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
      statusOptions: [],
      // 查询参数
      queryParams: {
        pageIndex: 1,
        pageSize: 10,
        id: undefined,
        cateId: undefined,
        name: undefined
      },
      // 表单参数
      form: {},
      // 表单校验
      rules: {
        cateId: [{ required: true, message: '分类编号不能为空', trigger: 'blur' }],
        name: [{ required: true, message: '标题不能为空', trigger: 'blur' }],
        status: [{ required: true, message: '状态不能为空', trigger: 'blur' }],
        remark: [{ required: true, message: '备注不能为空', trigger: 'blur' }],
        content: [{ required: true, message: '内容不能为空', trigger: 'blur' }]
      }
    }
  },
  created() {
    this.getList()
    this.getDicts('sys_status').then(response => {
      this.statusOptions = response.data
    })
  },
  methods: {
    /** 查询参数列表 */
    getList() {
      this.loading = true
      listContentArticle(this.addDateRange(this.queryParams, this.dateRange)).then(response => {
        this.tableList = response.data.list
        this.total = response.data.count
        this.loading = false
      })
    },
    /** 获取返回的信息 */
    getCategorySelectInfo(row) {
      this.form.cateId = row.id
      this.categoryName = row.name
      this.$refs['form'].validateField('cateId')
    },
    /** 显示选项框selectUserLevel */
    selectCategory() {
      this.$refs.category.show()
    },
    // 取消按钮
    cancel() {
      this.open = false
      this.reset()
    },
    // 表单重置
    reset() {
      this.categoryName = '请选择'
      this.form = {
        id: undefined,
        cateId: undefined,
        name: undefined,
        content: undefined,
        remark: undefined,
        status: undefined
      }
      this.resetForm('form')
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
      this.title = '添加文章'
      this.isEdit = false
    },
    // 修改按钮操作
    handleUpdate(row) {
      this.reset()
      getContentArticle(row.id).then(response => {
        this.form = response.data
        this.open = true
        this.title = '修改文章'
        this.isEdit = true
      })
    },
    /** 提交按钮 */
    submitForm: function() {
      this.$refs['form'].validate(valid => {
        if (valid) {
          if (this.form.id !== undefined) {
            updateContentArticle(this.form).then(response => {
              if (response.code === 200) {
                this.msgSuccess(response.msg)
                this.open = false
                this.getList()
              } else {
                this.msgError(response.msg)
              }
            })
          } else {
            addContentArticle(this.form).then(response => {
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
        return delContentArticle({ 'ids': ids })
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
        exportContentArticle(this.addDateRange(this.queryParams, this.dateRange)).then(response => {
          resolveBlob(response, '文章列表')
        })
      }).catch(() => {})
    }
  }
}
</script>
