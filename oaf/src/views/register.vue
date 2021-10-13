<style>
</style>
<template>
  <v-row class="fill-height" align="center" justify="center" style="background: #ebebeb">
    <v-col cols="12" sm="8" md="6" lg="4" xl="3">
      <v-card class="elevation-12 mx-5" style="opacity: 0.8">
        <v-row justify="center">
          <v-card class="elevation-1 mt-n7 primary" style="width: 80%">
            <v-card-actions>
              <v-row>
                <v-icon
                  style="position: absolute;left: 10px;top:19px;z-index: 1"
                  @click="$router.back()"
                  size="36"
                >mdi-arrow-left-circle
                </v-icon>
                <v-col cols="12" class="text-center">
                  <h1 class="display-2 ">注册</h1>
                </v-col>
              </v-row>
            </v-card-actions>
          </v-card>
        </v-row>
        <v-card-text class="text-center">
          <v-form ref="form">
            <v-text-field
              type="text"
              prepend-inner-icon="mdi-account-circle"
              v-model="form.username"
              label="账号"
              :rules="ruleInline.user"
              :counter="16"
            >
            </v-text-field>
            <v-text-field
              type="password"
              v-model="form.passwd"
              label="密码"
              prepend-inner-icon="mdi-lock"
              :rules="ruleInline.password"
              :counter="16"
            ></v-text-field>
            <v-text-field
              type="password"
              v-model="form.passwdCheck"
              label="密码"
              prepend-inner-icon="mdi-lock"
              :rules="ruleInline.passwordCheck"
              :counter="16"
              @keyup.enter="handleSubmit"
            ></v-text-field>
          </v-form>
        </v-card-text>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn type="primary" @click="handleSubmit">提交</v-btn>
          <v-btn @click="handleReset()">重置</v-btn>
        </v-card-actions>
      </v-card>
    </v-col>
  </v-row>
</template>

<script lang='ts'>
import {Component, Vue} from 'vue-property-decorator'

@Component({
  components: {}
})
export default class Register extends Vue {
  form = {
    passwd: '',
    passwdCheck: '',
    email: '',
    username: ''
  }

  ruleInline = {
    user: [
      (v: string) => !!v || 'required',
      (v: string) => (v && v.length >= 3 && v.length <= 16) || '长度要求3~16'
    ],
    password: [
      (v: string) => !!v || 'required',
      (v: string) => (v && v.length >= 6 && v.length <= 16) || '长度要求6~16'
    ],
    passwordCheck: [
      (v: string) => !!v || 'required',
      (v: string) => (v && v === this.form.passwd) || '密码不一致'
    ]
  }

  handleSubmit() {
    if (!this.$refs.form.validate()) {
      return
    }
    this.$api.user.register(this.form.username, this.form.passwd).Start(
      (data) => {
        // this.$message.success('注册成功!')
        this.$router.push({name: 'login'})
      },
      (data) => {
        if (data && data.code === '31011') {
          // this.$message.error('用户名重复')
        } else {
          // this.$message.error('注册失败')
        }
      }
    )
  }

  handleReset() {
    this.form.username = ''
    this.form.passwd = ''
    this.form.passwdCheck = ''
  }
}
</script>
