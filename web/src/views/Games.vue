<template>
  <section class="games-wrapper">
    <v-dialog v-model="isAddGameModal" width="500" @click:outside="clearForm">
      <v-card class="pa-6">
        <v-card-title>
          Добавить игру
        </v-card-title>
        <v-form @submit.prevent="sendGame">
          <v-text-field label="Имя игры" v-model="addGameForm.name"/>
          <v-text-field label="Описание игры" v-model="addGameForm.description"/>
          <v-text-field label="Ссылка на картинку" v-model="addGameForm.image"/>
          <v-btn type="submit" :disabled="!isSendButtonActive">
            Добавить
          </v-btn>
        </v-form>
      </v-card>
    </v-dialog>
    <v-card
      v-for="game in getGames"
      :key="game.id"
      width="300"
      height="500"
      class="game-card"
      @click="onGameClick(game)"
    >
      <v-card-title>
        {{ `${game.name}` }}
      </v-card-title>
      <v-card-subtitle>
        {{ `${prepareText(game.description)}` }}
      </v-card-subtitle>
      <img class="game-image" :src="'http://localhost:5000/games/' + game.id + '/image'" alt="game_img"/>
    </v-card>
    <v-card width="300" height="500" class="d-flex align-center justify-center" @click="isAddGameModal = true">
      <v-icon x-large>mdi-plus-thick</v-icon>
    </v-card>
  </section>
</template>

<script>

import { mapActions, mapGetters } from 'vuex'
export default {
  name: 'Games',
  data () {
    return {
      isAddGameModal: false,
      addGameForm: {
        name: '',
        description: '',
        image: ''
      }
    }
  },
  computed: {
    ...mapGetters('games', ['getGames']),
    isSendButtonActive () {
      const { name } = this.addGameForm
      return name
    }
  },
  methods: {
    ...mapActions('games', ['fetchGames', 'fetchCreateGame']),
    prepareText (text) {
      return text.length > 29 ? text.slice(0, 30) + '...' : text
    },
    sendGame () {
      this.fetchCreateGame(this.addGameForm)
        .then(() => {
          this.fetchGames()
          this.clearForm()
          this.isAddGameModal = false
        })
    },
    clearForm () {
      this.addGameForm = {
        name: '',
        description: '',
        image: ''
      }
    },
    onGameClick (game) {
      this.$router.push({ name: 'Game', params: { id: game.id } })
    }
  },
  mounted () {
    this.fetchGames()
  }
}
</script>
<style lang="scss" scoped>
.games-wrapper {
  display: flex;
  flex-wrap: wrap;
  gap: 20px;
}

.game-card {
  position: relative;
}

.game-image {
  width: 100%;
  height: 70%;
  position: absolute;
  bottom: 0;
  left: 0;
  right: 0
}
</style>
