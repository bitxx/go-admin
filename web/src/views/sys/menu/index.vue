<template>
  <BasicLayout>
    <template #wrapper>
      <el-card class="box-card">
        <el-form ref="queryForm" :model="queryParams" :inline="true" label-width="68px">
          <el-form-item label-width="100" label="菜单名称" prop="title">
            <el-input v-model="queryParams.title" placeholder="请输入菜单名称" clearable size="small" @keyup.enter.native="handleQuery" />
          </el-form-item>
          <el-form-item label="可见">
            <el-select v-model="queryParams.hidden" placeholder="可见状态" prop="hidden" clearable size="small">
              <el-option
                v-for="dict in hiddenOptions"
                :key="dict.dictValue"
                :label="dict.dictLabel"
                :value="dict.dictValue"
              />
            </el-select>
          </el-form-item>
          <el-form-item>
            <el-button type="primary" icon="el-icon-search" size="mini" @click="handleQuery">搜索</el-button>
            <el-button
              v-permisaction="['admin:sysMenu:add']"
              type="primary"
              icon="el-icon-plus"
              size="mini"
              @click="handleAdd"
            >新增</el-button>
          </el-form-item>
        </el-form>

        <el-table v-loading="loading" stripe border :data="menuList" row-key="id" :tree-props="{children: 'children', hasChildren: 'hasChildren'}">
          <el-table-column prop="title" label="菜单名称" :show-overflow-tooltip="true" width="220" />
          <el-table-column prop="icon" label="图标" align="center" width="100">
            <template slot-scope="scope">
              <svg-icon :icon-class="scope.row.icon" />
            </template>
          </el-table-column>
          <el-table-column width="60" align="center" prop="sort" label="排序" />
          <el-table-column width="260" align="center" prop="permission" label="权限标识" :show-overflow-tooltip="true">
            <template slot-scope="scope">
              <el-popover v-if="scope.row.sysApi.length>0" trigger="hover" placement="top">
                <el-table :data="scope.row.sysApi" border style="width: 100%">
                  <el-table-column prop="title" label="title" width="260px">
                    <template slot-scope="scope">
                      <span v-if="scope.row.apiType==='1' && scope.row.title!==''"><el-tag type="success">{{ '['+scope.row.apiType +'] '+ scope.row.title }}</el-tag></span>
                      <span v-if="scope.row.apiType==='2' && scope.row.title!==''"><el-tag type="warning">{{ '['+scope.row.apiType +'] '+ scope.row.title }}</el-tag></span>
                      <span v-if="scope.row.apiType!=='1' && scope.row.apiType!=='2' && scope.row.title!==''"><el-tag type="">{{ '['+scope.row.apiType +'] '+scope.row.title }}</el-tag></span>
                      <span v-if="scope.row.title===''"><el-tag type="danger">暂无</el-tag></span>
                    </template>
                  </el-table-column>
                  <el-table-column prop="path" label="path" width="270px">
                    <template slot-scope="scope">
                      <el-tag v-if="scope.row.action==='GET'">{{ scope.row.action }}</el-tag>
                      <el-tag v-if="scope.row.action==='POST'" type="success">{{ scope.row.action }}</el-tag>
                      <el-tag v-if="scope.row.action==='PUT'" type="warning">{{ scope.row.action }}</el-tag>
                      <el-tag v-if="scope.row.action==='DELETE'" type="danger">{{ scope.row.action }}</el-tag>
                      {{ scope.row.path }}
                    </template>
                  </el-table-column>
                </el-table>
                <div slot="reference" class="name-wrapper">
                  <span v-if="scope.row.permission===''">-</span>
                  <span v-else>{{ scope.row.permission }}</span>
                </div>
              </el-popover>
              <span v-else>
                <span v-if="scope.row.permission===''">-</span>
                <span v-else>{{ scope.row.permission }}</span>
              </span>
            </template>
          </el-table-column>
          <el-table-column width="100" label="菜单类型" align="center" prop="status" :formatter="menuTypeFormat">
            <template slot-scope="scope">
              {{ menuTypeFormat(scope.row) }}
            </template>
          </el-table-column>
          <el-table-column width="300" label="组建路径" align="center" prop="path" :show-overflow-tooltip="true">
            <template slot-scope="scope">
              <span v-if="scope.row.path===''">-</span>
              <span v-else>{{ scope.row.path }}</span>
            </template>
          </el-table-column>
          <el-table-column width="200" label="路由名称" align="center" prop="name" :show-overflow-tooltip="true">
            <template slot-scope="scope">
              <span v-if="scope.row.name===''">-</span>
              <span v-else>{{ scope.row.name }}</span>
            </template>
          </el-table-column>
          <el-table-column width="100" label="可见" align="center" prop="hidden" :formatter="hiddenFormat">
            <template slot-scope="scope">
              {{ hiddenFormat(scope.row) }}
            </template>
          </el-table-column>
          <el-table-column width="200" label="创建时间" align="center" prop="createdAt" :show-overflow-tooltip="true">
            <template slot-scope="scope">
              <span>{{ parseTime(scope.row.createdAt) }}</span>
            </template>
          </el-table-column>
          <el-table-column label="操作" fixed="right" align="center" class-name="small-padding fixed-width" width="180">
            <template slot-scope="scope">
              <el-button
                v-permisaction="['admin:sysMenu:edit']"
                size="mini"
                type="text"
                icon="el-icon-edit"
                @click="handleUpdate(scope.row)"
              >修改</el-button>
              <el-button
                v-permisaction="['admin:sysMenu:add']"
                size="mini"
                type="text"
                icon="el-icon-plus"
                @click="handleAdd(scope.row)"
              >新增</el-button>
              <el-button
                v-permisaction="['admin:sysMenu:remove']"
                size="mini"
                type="text"
                icon="el-icon-delete"
                @click="handleDelete(scope.row)"
              >删除</el-button>
            </template>
          </el-table-column>
        </el-table>

        <!-- 添加或修改菜单对话框 -->
        <el-drawer ref="drawer" :title="title" :visible.sync="open" direction="rtl" custom-class="demo-drawer" size="900px">
          <div class="demo-drawer__content">
            <el-form ref="form" :model="form" :rules="rules" label-position="top" label-width="106px">
              <el-row>
                <el-col :span="24">
                  <el-form-item prop="parentId">
                    <span slot="label">
                      上级菜单
                      <el-tooltip content="指当前菜单停靠的菜单归属" placement="top">
                        <i class="el-icon-question" />
                      </el-tooltip>
                    </span>
                    <treeselect v-model="form.parentId" :options="menuOptions" :normalizer="normalizer" :show-count="true" placeholder="选择上级菜单" />
                  </el-form-item>
                </el-col>
                <el-col :span="12">
                  <el-form-item prop="title">
                    <span slot="label">
                      菜单标题
                      <el-tooltip content="菜单位置显示的说明信息" placement="top">
                        <i class="el-icon-question" />
                      </el-tooltip>
                    </span>
                    <el-input v-model="form.title" placeholder="请输入菜单标题" />
                  </el-form-item>
                </el-col>

                <el-col :span="12">
                  <el-form-item prop="sort">
                    <span slot="label">
                      显示排序
                      <el-tooltip content="根据序号升序排列" placement="top">
                        <i class="el-icon-question" />
                      </el-tooltip>
                    </span>
                    <el-input-number v-model="form.sort" controls-position="right" :min="0" />
                  </el-form-item>
                </el-col>

                <el-col :span="24">
                  <el-form-item prop="menuType">
                    <span slot="label">
                      菜单类型
                      <el-tooltip content="包含目录：以及菜单或者菜单组，菜单：具体对应某一个页面，按钮：功能才做按钮；" placement="top">
                        <i class="el-icon-question" />
                      </el-tooltip>
                    </span>
                    <el-radio-group v-model="form.menuType">
                      <el-radio
                        v-for="dict in menuTypeOptions"
                        :key="dict.dictValue"
                        :label="dict.dictValue"
                      >{{ dict.dictLabel }}</el-radio>
                    </el-radio-group>
                  </el-form-item>
                </el-col>
                <el-col :span="24">
                  <el-form-item label="菜单图标">
                    <el-popover placement="bottom-start" width="460" trigger="click" @show="$refs['iconSelect'].reset()">
                      <IconSelect ref="iconSelect" @selected="selected" />
                      <el-input slot="reference" v-model="form.icon" placeholder="点击选择图标" readonly>
                        <svg-icon
                          v-if="form.icon"
                          slot="prefix"
                          :icon-class="form.icon"
                          class="el-input__icon"
                          style="height: 32px;width: 16px;"
                        />
                        <i v-else slot="prefix" class="el-icon-search el-input__icon" />
                      </el-input>
                    </el-popover>
                  </el-form-item>
                </el-col>
                <el-col :span="12">
                  <el-form-item v-if="form.menuType === '2'" prop="name">
                    <span slot="label">
                      路由名称
                      <el-tooltip content="需要和页面name保持一致，对应页面即可选择缓存" placement="top">
                        <i class="el-icon-question" />
                      </el-tooltip>
                    </span>
                    <el-input v-model="form.name" placeholder="请输入路由名称" />
                  </el-form-item>
                </el-col>

                <el-col v-if="form.menuType === '1' || form.menuType === '2'" :span="12">
                  <el-form-item prop="component">
                    <span slot="label">
                      组件路径
                      <el-tooltip content="菜单对应的具体vue页面文件路径views的下级路径/admin/sys-api/index；目录类型：填写Layout，如何有二级目录请参照日志目录填写；" placement="top">
                        <i class="el-icon-question" />
                      </el-tooltip>
                    </span>
                    <el-input v-model="form.component" placeholder="请输入组件路径" />
                  </el-form-item>
                </el-col>

                <el-col :span="12">
                  <el-form-item v-if="form.menuType === '1' || form.menuType === '2'">
                    <span slot="label">
                      是否外链
                      <el-tooltip content="可以通过iframe打开指定地址" placement="top">
                        <i class="el-icon-question" />
                      </el-tooltip>
                    </span>
                    <el-radio-group v-model="form.isFrame">
                      <el-radio
                        v-for="dict in isFrameOptions"
                        :key="dict.dictValue"
                        :label="dict.dictValue"
                      >{{ dict.dictLabel }}</el-radio>
                    </el-radio-group>
                  </el-form-item>
                </el-col>

                <el-col :span="12">
                  <el-form-item v-if="form.menuType !== '3'" prop="path">
                    <span slot="label">
                      路由地址
                      <el-tooltip content="访问此页面自定义的url地址，建议/开头书写，例如 /app-name/menu-name" placement="top">
                        <i class="el-icon-question" />
                      </el-tooltip>
                    </span>
                    <el-input v-model="form.path" placeholder="请输入路由地址" />
                  </el-form-item>
                </el-col>

                <el-col :span="12">
                  <el-form-item v-if="form.menuType === '3'">
                    <span slot="label">
                      权限标识
                      <el-tooltip content="前端权限控制按钮是否显示" placement="top">
                        <i class="el-icon-question" />
                      </el-tooltip>
                    </span>
                    <el-input v-model="form.permission" placeholder="请权限标识" maxlength="50" />
                  </el-form-item>
                </el-col>
                <el-col :span="12">
                  <el-form-item v-if="form.menuType === '1' || form.menuType === '2'">
                    <span slot="label">
                      菜单状态
                      <el-tooltip content="需要显示在菜单列表的菜单设置为显示，否则设置为隐藏" placement="top">
                        <i class="el-icon-question" />
                      </el-tooltip>
                    </span>
                    <el-radio-group v-model="form.hidden">
                      <el-radio
                        v-for="dict in hiddenOptions"
                        :key="dict.dictValue"
                        :label="dict.dictValue"
                      >{{ dict.dictLabel }}</el-radio>
                    </el-radio-group>
                  </el-form-item>
                </el-col>
                <el-col :span="24">
                  <el-form-item v-if="form.menuType === '2' || form.menuType === '3'">
                    <span slot="label">
                      api权限
                      <el-tooltip content="配置在这个才做上需要使用到的接口，否则在设置用户角色时，接口将无权访问。" placement="top">
                        <i class="el-icon-question" />
                      </el-tooltip>
                    </span>
                    <el-transfer
                      v-model="form.apis"
                      style="text-align: left; display: inline-block"
                      filterable
                      :props="{
                        key: 'id',
                        label: 'title'
                      }"
                      :titles="['未授权', '已授权']"
                      :button-texts="['收回', '授权 ']"
                      :format="{
                        noChecked: '${total}',
                        hasChecked: '${checked}/${total}'
                      }"
                      class="panel"
                      :data="sysapiList"
                      @change="handleChange"
                    >
                      <span slot-scope="{ option }">{{ option.title }}</span>
                    </el-transfer>
                  </el-form-item>
                </el-col>
              </el-row>
            </el-form>
            <div class="demo-drawer__footer">
              <el-button type="primary" @click="submitForm">确 定</el-button>
              <el-button @click="cancel">取 消</el-button>
            </div>
          </div>
        </el-drawer>
      </el-card>
    </template>
  </BasicLayout>
</template>

<script>
import { listMenu, getMenu, delMenu, addMenu, updateMenu } from '@/api/sys/menu'
import { listSysApi } from '@/api/sys/api'

import Treeselect from '@riophae/vue-treeselect'
import '@riophae/vue-treeselect/dist/vue-treeselect.css'
import IconSelect from '@/components/IconSelect/index.vue'

export default {
  name: 'SysMenu',
  components: { Treeselect, IconSelect },
  data() {
    return {
      // 遮罩层
      loading: true,
      // 菜单表格树数据
      menuList: [],
      sysapiList: [],
      // 菜单树选项
      menuOptions: [],
      // 弹出层标题
      title: '',
      // 是否显示弹出层
      open: false,
      // 菜单状态数据字典
      hiddenOptions: [],
      isFrameOptions: [],
      menuTypeOptions: [],
      // 查询参数
      queryParams: {
        title: undefined,
        hidden: undefined
      },
      // 表单参数
      form: {
        apis: [],
        sysApi: []
      },
      // 表单校验
      rules: {
        title: [{ required: true, message: '菜单标题不能为空', trigger: 'blur' }],
        sort: [{ required: true, message: '菜单顺序不能为空', trigger: 'blur' }]
      }
    }
  },
  created() {
    this.getList()

    this.getApiList()
    this.getDicts('sys_menu_show_hide').then(response => {
      this.hiddenOptions = response.data
    })
    this.getDicts('sys_yes_no').then(response => {
      this.isFrameOptions = response.data
    })
    this.getDicts('sys_menu_type').then(response => {
      this.menuTypeOptions = response.data
    })
  },
  methods: {
    menuTypeFormat(row) {
      return this.selectDictLabel(this.menuTypeOptions, row.menuType)
    },
    handleChange(value, direction, movedKeys) {
      const list = this.form.sysApi
      this.form.apis = value
      if (direction === 'right') {
        for (let x = 0; x < movedKeys.length; x++) {
          for (let index = 0; index < this.sysapiList.length; index++) {
            const element = this.sysapiList[index]
            if (element.id === movedKeys[x]) {
              list.push(element)
              break
            }
          }
        }
        this.form.sysApi = list
      } else if (direction === 'left') {
        const l = []
        for (let index = 0; index < movedKeys.length; index++) {
          const element = movedKeys[index]
          for (let x = 0; x < list.length; x++) {
            const e = list[x]
            if (element !== e.id) {
              l.push()
              break
            }
          }
        }
        this.form.sysApi = l
      }
      // this.setApis(this.form.SysApi)
    },
    getApiList() {
      this.loading = true
      // 取出所有接口
      listSysApi({ 'pageSize': 10000, 'apiTypes': ['2', '3'] }).then(response => {
        this.sysapiList = response.data.list
        this.loading = false
      }
      )
    },
    // 选择图标
    selected(name) {
      this.form.icon = name
    },
    /** 查询菜单列表 */
    getList() {
      this.loading = true
      listMenu(this.queryParams).then(response => {
        this.menuList = response.data.list
        this.loading = false
      })
    },
    /** 转换菜单数据结构 */
    normalizer(node) {
      if (node.children && !node.children.length) {
        delete node.children
      }
      return {
        id: node.id,
        label: node.title,
        children: node.children
      }
    },
    /** 查询菜单下拉树结构 */
    getTreeselect() {
      listMenu().then(response => {
        this.menuOptions = []
        const menu = { id: 0, title: '主类目', children: [] }
        menu.children = response.data.list
        this.menuOptions.push(menu)
      })
    },
    // 菜单显示状态字典翻译
    hiddenFormat(row) {
      return this.selectDictLabel(this.hiddenOptions, row.hidden)
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
        parentId: 0,
        name: undefined,
        icon: undefined,
        menuType: undefined,
        apis: [],
        sort: 0,
        action: undefined,
        isFrame: undefined,
        hidden: undefined
      }
      this.resetForm('form')
    },
    /** 搜索按钮操作 */
    handleQuery() {
      this.getList()
    },
    /** 新增按钮操作 */
    handleAdd(row) {
      this.reset()
      this.getTreeselect()
      if (row != null) {
        this.form.parentId = row.id
      }
      this.open = true
      this.title = '添加菜单'
    },
    /** 修改按钮操作 */
    handleUpdate(row) {
      this.reset()
      this.getTreeselect()
      getMenu(row.id).then(response => {
        this.form = response.data
        this.open = true
        this.title = '修改菜单'
      })
    },
    setApis(apiArray) {
      var l = []
      for (var index = 0; index < apiArray.length; index++) {
        const element = apiArray[index]
        l.push(element.id)
      }
      this.form.apis = l
    },
    /** 提交按钮 */
    submitForm: function() {
      this.$refs['form'].validate(valid => {
        if (valid) {
          if (this.form.id !== undefined) {
            updateMenu(this.form, this.form.id).then(response => {
              if (response.code === 200) {
                this.msgSuccess(response.msg)
                this.open = false
                this.getList()
              } else {
                this.msgError(response.msg)
              }
            })
          } else {
            addMenu(this.form).then(response => {
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
      this.$confirm('是否确认删除名称为"' + row.title + '"的数据项?', '警告', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(function() {
        const ids = [row.id]
        return delMenu({ 'ids': ids })
      }).then((response) => {
        if (response.code === 200) {
          this.msgSuccess(response.msg)
          this.open = false
          this.getList()
        } else {
          this.msgError(response.msg)
        }
      }).catch(function() {})
    }
  }
}
</script>
<style lang="css">
.panel .el-transfer__buttons{
  width: 150px;
}
.panel .el-transfer__buttons .el-button + .el-button{
  margin-left:0;
}
.panel .el-transfer-panel{
  width: 300px;
}

.el-col {
padding: 0 5px;
}
.el-drawer__header{
margin-bottom: 0;
}
</style>
