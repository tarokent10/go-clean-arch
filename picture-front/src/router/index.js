import Vue from 'vue'
import VueRouter from 'vue-router'
import TitlePage from '@/pages/TitlePage.vue'
import CardGame from '@/pages/CardGame.vue'
import RegistImage from '@/pages/RegistImage.vue'
import Images from '@/pages/Images.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'titlePage', // タイトルページできたら修正
    component: TitlePage
  },
  {
    path: '/cardgame',
    name: 'cardgame',
    component: CardGame
  },
  {
    path: '/registImage',
    name: 'registImage',
    component: RegistImage
  },
  {
    path: '/images',
    name: 'images',
    component: Images
  }
]

const router = new VueRouter({
  routes
})

export default router
