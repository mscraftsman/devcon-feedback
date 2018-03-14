<template>
  <div class="page page-feedback">
    <div class="logo-wrapper">
      <router-link :to="{ name: 'login' }" >
        <img src="mscc-logo.png" alt="">
      </router-link>
    </div>
    <div class="user-info">
      <span class="title">Logged in as</span>
      <span>{{ getName }}</span>
      <span class="title">currently rating</span>
      <span>{{ getSessionCurrent.questionCurrent }}</span>
    </div>

    <div class="questions-wrapper" >
        <div class="virer-mam">
            <div class="devirer-mam">
                <transition name="slide-fade"
                    v-for="(question, index) in getQuestionText" 
                    :key="index">
                    <div class="question-holder" 
                        v-show="getSessionCurrent.questionCurrent == index">
                        <div class="question">{{ question }}</div>
                        <div v-if="index === 0" class="comment-wrapper">
                            <div class="reaction-wrapper">
                                <Reactions />
                            </div>
                        </div>
                        <div v-if="index === 1" class="comment-wrapper">
                            <div class="reaction-wrapper">
                                <Reactions />
                            </div>
                        </div>
                        <div v-else-if="index === 2" class="comment-wrapper">
                            <div class="reaction-wrapper">
                                <Reactions type="yesno"/>
                            </div>
                        </div>
                        <div v-else-if="index === 3" class="comment-wrapper">
                            <textarea class="textbox" placeholder="We would be grateful if you could leave some constructive critism."></textarea>
                        </div>
                    </div>
                </transition>
                <div class="button-next" @click="$store.commit('NEXT_QUESTION')">Next</div>
            </div>
        </div>
    </div>

    <div class="footer">
      Developer Conference 2018
    </div>


  </div>
</template>

<script>
import Reactions from 'components/Reactions'
import {mapGetters} from 'vuex'

export default {
    methods: {
        next: function () {
            console.log('next')
        }
    },
    computed: {
        ...mapGetters([
            'getName',
            'getSessionCurrent',
            'getQuestionText'
            ])
    },
    components: {
        Reactions
    }
}
</script>

<style lang="scss" scoped>

/* Enter and leave animations can use different */
/* durations and timing functions.              */
.slide-fade-enter-active {
  transition: all .3s ease;
}
.slide-fade-leave-active {
  transition: all .8s cubic-bezier(1.0, 0.5, 0.8, 1.0);
}
.slide-fade-enter{
//   transform: scale(0.5);
  opacity: 0;
}
.slide-fade-leave-to
/* .slide-fade-leave-active below version 2.1.8 */ {
  transform: scale(2);
  opacity: 0;
}

.logo-wrapper {
  img {
    padding: 15px;
    display: block;
    height: 100%;
    width: auto;
  }
}

.user-info {
  padding: 18px;
  span {
    display:block;
    text-align: right;

    &.title {
      text-transform: uppercase;
      font-size: 12px;
      color:white;

      &:not(:first-child) {
        padding-top: 10px ;
      }
    }
  }
  font-size: 14px;
}

.page-feedback {
  display: grid;
  grid-template-areas: 
    "logo userinfo"
    "questions questions"
    "footer footer";
  grid-template-columns: 100px 1fr;
  grid-template-rows: 10vh 80vh 10vh;
}

.questions-wrapper {
  grid-area: questions;
  align-items: center;
  justify-content: center;
  align-content: center;
  display:grid;
  overflow: hidden;

  .virer-mam {
      transform: rotate(-10deg);
      background: white;
      width: 120vw;
  }

  .devirer-mam {
      transform: rotate(10deg);
      padding: 50px 10vw;
      display:grid;
      grid-template-areas: "here";
      grid-template-columns: 1;
      grid-template-rows: 40vh;
    //   width: 100vw;

     .question-holder {
         grid-area: here;
     }
  }

  .question {
      font-weight: 600;
      font-family: "Roboto", sans-serif;
      font-size: 20px;
      padding: 0 20px;
  }
}

.button-next {
    background: red;
}

.textbox {
    // border: 1px solid green;
    width: 80%;
    height: 20vh;
    font-size: 20px;
    padding: 10px;
    margin-top: 10px;
    box-shadow: 0 0 4px #999;
    resize: disable;
}

.footer {
  grid-area: footer;
  color:white;
  text-align: center;
  font-size: 13px;
  align-self: center;
}
</style>

