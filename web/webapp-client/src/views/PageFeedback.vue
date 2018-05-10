<template>
  <div class="page page-feedback">
    <div class="info" v-if="session">
      <span class="title">currently rating {{ session.title }} x</span>
      <!-- <span>{{ getSessionCurrent.questionCurrent }}</span> -->
    </div>

    <div class="questions-wrapper">
      <div class="virer-mam">
        <div class="devirer-mam">
          <transition name="slide-fade" v-for="(question, index) in questions" :key="index">
            <div class="question-holder" v-show="getSessionCurrent.questionCurrent == index">
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
                  <Reactions type="yesno" />
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

    <!-- <div class="footer">
      Developer Conference 2018
    </div> -->

  </div>
</template>

<script>
import Reactions from "../components/Reactions";
import { mapGetters } from "vuex";

export default {
  props: ["id"],
  data() {
    return {
      reaction: {
        session_id: this.id,
        visitor_id: "userid",
        reaction_1: "",
        reaction_2: "",
        reaction_3: "",
        reaction_4: "",
        created_at: ""
      },
      questions: [
        "How would you rate the speaker?",
        "How would you rate the session content?",
        "Did you learn something new?",
        "Remarks for speaker or organisation?"
      ]
    };
  },
  methods: {
    next: function() {
      console.log("next");
    }
  },
  computed: {
    ...mapGetters(["getName", "getSessionCurrent"]),
    ...mapGetters({
      sessions: "getSessions",
      speakers: "getSpeakers"
    }),
    session: function() {
      let sessions = this.sessions
        .map(groups => groups.sessions)
        .reduce(function(acc, curr) {
          return [...acc, ...curr];
        }, []);
      let session = _.filter(sessions, { id: this.id })[0];
      // console.log(session);
      return session;
    }
  },
  components: {
    Reactions
  }
};
</script>

<style lang="scss" scoped>
/* Enter and leave animations can use different */
/* durations and timing functions.              */
.slide-fade-enter-active {
  transition: all 0.3s ease;
}
.slide-fade-leave-active {
  transition: all 0.8s cubic-bezier(1, 0.5, 0.8, 1);
}
.slide-fade-enter {
  //   transform: scale(0.5);
  opacity: 0;
}
.slide-fade-leave-to {
  transform: scale(2);
  opacity: 0;
}

.info {
  grid-area: userinfo;
  padding: 18px;
  span {
    display: block;
    text-align: right;

    &.title {
      text-transform: uppercase;
      font-size: 12px;
      color: white;

      &:not(:first-child) {
        padding-top: 10px;
      }
    }
  }
  font-size: 14px;
}

.page-feedback {
  display: grid;
  grid-template-areas:
    "userinfo userinfo"
    "questions questions";
  // "footer footer";
  grid-template-columns: 100px 1fr;
  grid-template-rows: 10vh 80vh;
}

.questions-wrapper {
  grid-area: questions;
  align-items: center;
  justify-content: center;
  align-content: center;
  display: grid;
  overflow: hidden;

  .virer-mam {
    transform: rotate(-10deg);
    background: white;
    width: 120vw;
  }

  .devirer-mam {
    transform: rotate(10deg);
    padding: 50px 10vw;
    display: grid;
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
  color: white;
  text-align: center;
  font-size: 13px;
  align-self: center;
}
</style>

