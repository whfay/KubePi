<template>
  <layout-content header="Services">
    <div style="float: left">
      <el-button v-has-permissions="{scope:'namespace',apiGroup:'',resource:'services',verb:'create'}"
                  type="primary" size="small"
                  @click="yamlCreate">
        YAML
      </el-button>
      <el-button type="primary" size="small" @click="onCreate"
                  v-has-permissions="{scope:'namespace',apiGroup:'',resource:'services',verb:'create'}">
        {{ $t("commons.button.create") }}
      </el-button>
      <el-button type="primary" size="small" :disabled="selects.length===0" @click="onDelete()"
                  v-has-permissions="{scope:'namespace',apiGroup:'',resource:'services',verb:'delete'}">
        {{ $t("commons.button.delete") }}
      </el-button>
    </div>
    <complex-table :data="data" :selects.sync="selects" @search="search" v-loading="loading"
                   :pagination-config="paginationConfig" :search-config="searchConfig">
      <el-table-column type="selection" fix></el-table-column>
      <el-table-column :label="$t('commons.table.name')" prop="metadata.name" show-overflow-tooltip>
        <template v-slot:default="{row}">
          <span class="span-link" @click="openDetail(row)">{{ row.metadata.name }}</span>
        </template>
      </el-table-column>
      <el-table-column :label="$t('business.namespace.namespace')" prop="metadata.namespace">
        <template v-slot:default="{row}">
          {{ row.metadata.namespace }}
        </template>
      </el-table-column>
      <el-table-column :label="$t('business.configuration.type')" prop="spec.type">
        <template v-slot:default="{row}">
          {{ row.spec.type }}
        </template>
      </el-table-column>
      <el-table-column label="ClusterIP" prop="spec.clusterIP" min-width="100px" show-overflow-tooltip>
        <template v-slot:default="{row}">
          {{ row.spec.clusterIP || "" }}
        </template>
      </el-table-column>
      <el-table-column :label="$t('business.network.target_port')" min-width="120px" show-overflow-tooltip>
        <template v-slot:default="{row}">
          <div v-for="(value,key,index) in row.spec.ports" v-bind:key="index" type="info" size="mini">
            <span style="font-size: 12px" v-if="row.spec.type ==='NodePort'">
              {{ value.port }}:{{ value.nodePort }}/{{ value.protocol }}
            </span>
            <span style="font-size: 12px" v-if="row.spec.type !=='NodePort' && value.port === value.targetPort">
              {{ value.port }}/{{ value.protocol }}
            </span>
            <span style="font-size: 12px" v-if="row.spec.type !=='NodePort' && value.port !== value.targetPort">
              {{ value.port }}:{{ value.targetPort }}/{{ value.protocol }}
            </span>
            <br>
          </div>
        </template>
      </el-table-column>
      <el-table-column :label="$t('business.network.selector')" prop="spec.selector" min-width="200px">
        <template v-slot:default="{row}">
          <el-tag v-for="(value,key,index) in row.spec.selector" v-bind:key="index" type="info" size="mini">
            {{ key }}={{ value }}
          </el-tag>
        </template>
      </el-table-column>
      <el-table-column :label="$t('commons.table.created_time')" prop="metadata.creationTimestamp" fix>
        <template v-slot:default="{row}">
          {{ row.metadata.creationTimestamp | age }}
        </template>
      </el-table-column>
      <ko-table-operations :buttons="buttons" :label="$t('commons.table.action')"></ko-table-operations>
    </complex-table>
  </layout-content>
</template>

<script>
import LayoutContent from "@/components/layout/LayoutContent"
import {downloadYaml} from "@/utils/actions"
import {deleteService, getService, listServices} from "@/api/services"
import ComplexTable from "@/components/complex-table"
import KoTableOperations from "@/components/ko-table-operations"
import {checkPermissions} from "@/utils/permission"

export default {
  name: "Services",
  components: { LayoutContent, ComplexTable, KoTableOperations },
  data () {
    return {
      data: [],
      selects: [],
      cluster: "",
      loading: false,
      buttons: [
        {
          label: this.$t("commons.button.edit"),
          icon: "el-icon-edit",
          click: (row) => {
            this.$router.push({
              name: "ServiceEdit",
              params: { namespace: row.metadata.namespace, name: row.metadata.name },
              query: { yamlShow: false }
            })
          },
          disabled: () => {
            return !checkPermissions({ scope: "namespace", apiGroup: "", resource: "services", verb: "update" })
          }
        },
        {
          label: this.$t("commons.button.edit_yaml"),
          icon: "el-icon-edit",
          click: (row) => {
            this.$router.push({
              name: "ServiceEdit",
              params: { name: row.metadata.name, namespace: row.metadata.namespace },
              query: { yamlShow: true }
            })
          },
          disabled: () => {
            return !checkPermissions({ scope: "namespace", apiGroup: "", resource: "services", verb: "update" })
          }
        },
        {
          label: this.$t("commons.button.download_yaml"),
          icon: "el-icon-download",
          click: (row) => {
            downloadYaml(row.metadata.name + ".yml",getService(this.cluster,row.metadata.namespace,row.metadata.name))
          }
        },
        {
          label: this.$t("commons.button.delete"),
          icon: "el-icon-delete",
          click: (row) => {
            this.onDelete(row)
          },
          disabled: () => {
            return !checkPermissions({ scope: "namespace", apiGroup: "", resource: "services", verb: "delete" })
          }
        },
      ],
      paginationConfig: {
        currentPage: 1,
        pageSize: 10,
        total: 0,
      },
      searchConfig: {
        keywords: ""
      }
    }
  },
  methods: {
    search (resetPage) {
      this.loading = true
      if (resetPage) {
        this.paginationConfig.currentPage = 1
      }
      listServices(this.cluster, true, this.searchConfig.keywords, this.paginationConfig.currentPage, this.paginationConfig.pageSize).then(res => {
        this.data = res.items
        this.loading = false
        this.paginationConfig.total = res.total
      })
    },
    onCreate () {
      this.$router.push({
        name: "ServiceCreate", query: { yamlShow: false }
      })
    },
    yamlCreate () {
      this.$router.push({
        name: "ServiceCreateYaml", query: { type: "services" }
      })
    },
    onDelete (row) {
      this.$confirm(
        this.$t("commons.confirm_message.delete"),
        this.$t("commons.message_box.prompt"), {
          confirmButtonText: this.$t("commons.button.confirm"),
          cancelButtonText: this.$t("commons.button.cancel"),
          type: "warning",
        }).then(() => {
        this.ps = []
        if (row) {
          this.ps.push(deleteService(this.cluster, row.metadata.namespace, row.metadata.name))
        } else {
          if (this.selects.length > 0) {
            for (const select of this.selects) {
              this.ps.push(deleteService(this.cluster, select.metadata.namespace, select.metadata.name))
            }
          }
        }
        if (this.ps.length !== 0) {
          Promise.all(this.ps)
            .then(() => {
              this.search(true)
              this.$message({
                type: "success",
                message: this.$t("commons.msg.delete_success"),
              })
            })
            .catch(() => {
              this.search(true)
            })
        }
      })
    },
    openDetail (row) {
      this.$router.push({
        name: "ServiceDetail",
        params: { name: row.metadata.name, namespace: row.metadata.namespace },
        query: { yamlShow: false }
      })
    },
  },
  created () {
    this.cluster = this.$route.query.cluster
    this.search()
  }
}
</script>

<style scoped>

</style>
