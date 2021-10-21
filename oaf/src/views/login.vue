<style>
</style>
<template>
  <v-row align="center" class="fill-height" justify="center" style="background: #ebebeb">
    <v-col cols="12" sm="8" md="6" lg="4" xl="3">
      <v-card class="elevation-12 mx-5" style="opacity: 0.8">
        <v-row justify="center">
          <v-col cols="10">
            <v-card class="elevation-1 mt-n12 primary theme--dark">
              <v-card-text class="text-center">
                <h1 class="display-2 font-weight-bold mb-2">Login</h1>
                <v-tooltip left>
                  <template v-slot:activator="{ on }">
                  <v-btn icon large v-on="on">
                    <v-icon>mdi-cellphone</v-icon>
                  </v-btn>
                  </template>
                  <span style="font-family:'Noto Sans Armenian'">手机登录</span>
                </v-tooltip>
                <v-tooltip right>
                  <template v-slot:activator="{ on }">
                  <v-btn icon large v-on="on">
                    <v-icon>mdi-barcode</v-icon>
                  </v-btn>
                  </template>
                  <span>授权码登录</span>
                </v-tooltip>
              </v-card-text>
            </v-card>
          </v-col>
        </v-row>
        <v-card-text>
          <v-form ref="form">
            <v-text-field
              v-model="formInline.user"
              :counter="16"
              :rules="ruleInline.user"
              label="账号"
              required
              prepend-inner-icon="mdi-account-circle"
            ></v-text-field>
            <v-text-field
              v-model="formInline.password"
              type="password"
              :counter="16"
              :rules="ruleInline.password"
              label="密码"
              prepend-inner-icon="mdi-lock"
              @keyup.enter="handleSubmit"
              required
            ></v-text-field>
          </v-form>
        </v-card-text>
        <v-card-actions>
          <v-spacer/>
          <v-btn type="primary" @click="handleSubmit">登录</v-btn>
          <router-link :to="{name: 'register', query:$route.query, params: $route.params}"
                       style="text-decoration: none;">
            <v-btn type="primary" style="margin-left:8px">注册</v-btn>
          </router-link>
        </v-card-actions>
      </v-card>
    </v-col>
  </v-row>
</template>

<script lang='ts'>
import {Component, Vue} from 'vue-property-decorator'
import util from '@/libs/util'

@Component({
  components: {}
})
export default class Login extends Vue {
  formInline = {
    user: '',
    password: ''
  }

  ruleInline = {
    user: [
      (v: string) => !!v || 'required',
      (v: string) => (v && v.length >= 3 && v.length <= 16) || '长度要求3~16'
    ],
    password: [
      (v: string) => !!v || 'required',
      (v: string) => (v && v.length >= 6 && v.length <= 16) || '长度要求6~16'
    ]
  }

  get app_uuid() {
    return this.$route.params.uuid || this.$store.state.oauuid
  }

  handleSubmit() {
    // eslint-disable-next-line
    // @ts-ignore
    if (!this.$refs.form.validate()) {
      return
    }
    this.$api.user.login(this.formInline.user, this.formInline.password, this.app_uuid).Start(
      data => {
        console.log(data)
        if (util.checkLogin()) {
          // this.$message.success('登录成功')
          // EventBus.$emit('login', true)
          this.$nextTick(() => {
            if (this.$route.query.redirect) {
              window.location.href = this.$route.query.redirect as string
            }
            this.$router.push({name: 'home'})
          })
        } else {
          // this.$message.error('用户名或密码错误')
        }
      },
      () => {
        // this.$message.error('网络错误！')
      }
    )
  }
}
</script>
