<template>
  <div id="app">
    <v-app id="inspire">
      <v-app-bar
        app
        color="cyan"
        dark
      >
        <v-spacer />
        <v-toolbar-title>Menu</v-toolbar-title>
        <v-app-bar-nav-icon @click.stop="drawer = !drawer"></v-app-bar-nav-icon>
          <v-btn
            color="primary"
            large
            @click="logout"
          >
          ログアウト
        </v-btn>
      </v-app-bar>
      <v-navigation-drawer
        v-model="drawer"
        absolute
        temporary
      >
        <v-list nav dense>
          <v-list-item-group
            active-class="deep-purple--text text--accent-4"
          >
            <v-list-item @click="toRegist">
              <v-list-item-title>従業員を登録する</v-list-item-title>
            </v-list-item>
            <v-list-item @click="toImages">
              <v-list-item-title>従業員の一覧を見る</v-list-item-title>
            </v-list-item>
          </v-list-item-group>
        </v-list>
      </v-navigation-drawer>
      <v-main>
        <router-view />
      </v-main>

      <v-footer
        color="cyan"
        app
      >
        <v-spacer />

        <span class="white--text">&copy; {{ new Date().getFullYear() }}</span>
      </v-footer>
    </v-app>
  </div>
</template>

<script>
import axios from 'axios'
export default {
  name: 'App',
  components: {},
  data: function () {
    return {
      drawer: false
    }
  },
  methods: {
    toRegist: function () {
      this.$router.push('/registImage')
    },
    toImages: function () {
      this.$router.push('/images')
    },
    logout: async function () {
      try {
        const instance = axios.create({})
        await instance.post('/v1/auth/logout/', '{}')
        this.$router.push('/')
      } catch (error) {
        console.log(error)
      }
    }
  }
}
</script>
