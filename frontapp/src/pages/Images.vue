<template>
  <v-container class="fill-height" fluid>
    <v-row>
      <!-- <transition-group name="flip-list" tag="div" class="container"> -->
      <v-col v-for="(image, idx) in images" :key=idx>
        <v-card width="200px" height="230px" elevation="5" >
          <v-img
            v-bind:src="image.uri"
            class="white--text align-end"
            gradient="to bottom, rgba(0,0,0,.1), rgba(0,0,0,.5)"
          >
          </v-img>
        </v-card>
      </v-col>
      <!-- </transition-group> -->
    </v-row>

  </v-container>
</template>
<script>
import axios from 'axios'
export default {
  name: 'images',
  data: function () {
    return {
      images: []
    }
  },
  methods: {
    get: function () {
      const is = this.images
      axios.get('/v1/images/', {
        withCredentials: true
      }).then(function (res) {
        res.data.forEach(e => {
          const d = 'data:image/png;base64,' + e.data
          is.push({ uri: d })
        })
        console.log('post sucessed')
      })
        .catch(function (error) {
          console.log(error)
        })
    }
  },
  created: function () {
    this.get()
  }
}
</script>
<style scoped>
</style>
