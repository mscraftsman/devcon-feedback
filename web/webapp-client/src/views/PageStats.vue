<template>
    <div class="page">
        <h1>Statistics</h1>

        <div class="stats-wrapper">

            <h2>How would you rate the speaker?</h2>
            <div class="chart-wrapper">
                <div class="speakers" v-if="stats.sessions">
                    <div class="speaker" v-for="session in _.orderBy(stats.sessions, 'reaction_1', ['desc']).slice(0, 10)" :key="session.session_id">
                        <router-link :to="{ name: 'session',  params: { id: session.session_id }}" class="name">{{ session.session_id }}</router-link>
                        <div class="graph rectangle">
                            <div class="bar" :style="{ width: (parseInt(session.reaction_1)/max_reaction_1) * 100 + '%' }">
                                <span class="bar-label" v-if="session.speaker"> {{ session.speakers }}</span>
                            </div>
                        </div>
                        <div class="value">
                            {{ session.reaction_1 }}
                        </div>
                    </div>
                </div>
            </div>
        </div>
        <div class="stats-wrapper">
            <h2>How would you rate the session content?</h2>
            <div class="chart-wrapper">
                <div class="speakers" v-if="stats.sessions">
                    <div class="speaker" v-for="session in _.orderBy(stats.sessions, 'reaction_2', ['desc']).slice(0, 10)" :key="session.session_id">
                        <router-link :to="{ name: 'session',  params: { id: session.session_id }}" class="name"> {{ session.session_id }}</router-link>
                        <div class="graph rectangle">
                            <div class="bar" :style="{ width: (parseInt(session.reaction_2)/max_reaction_2) * 100 + '%' }">
                                <span class="bar-label" v-if="session.speaker">{{ session.speakers }}</span>
                            </div>
                        </div>
                        <div class="value">
                            {{ session.reaction_2 }}
                        </div>
                    </div>
                </div>
            </div>
        </div>
        <div class="stats-wrapper">
            <h2>Top Voters</h2>
            <div class="chart-wrapper">
                <div class="speakers" v-if="stats.voters">
                    <div class="speaker" v-for="session in stats.voters.slice(0, 5)" :key="session.session_id">
                        <a target="_blank" :href="'https://www.meetup.com/MauritiusSoftwareCraftsmanshipCommunity/members/'+ session.id" class="name">
                            <!-- {{ session.id }} -->
                            <img class="photo" :src="session.photo" alt="">
                        </a>

                        <div class="graph rectangle">
                            <div class="bar" :style="{ width: (parseInt(session.votes)/max_votes) * 100 + '%' }">
                                <span v-if="session.name">{{ session.name }}</span>
                            </div>
                        </div>
                        <div class="value">
                            {{ session.votes }}
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
import { mapGetters, mapActions } from "vuex";
export default {
  data() {
    return {};
  },
  computed: {
    ...mapGetters({
      stats: "getStats"
    }),
    max_reaction_1() {
      let total = _.map(this.stats.sessions, "reaction_1");
      let max = _.max(total);
      console.log(max);
      return max;
    },
    max_reaction_2() {
      let total = _.map(this.stats.sessions, "reaction_2");
      let max = _.max(total);
      console.log(max);
      return max;
    },
    max_votes() {
      let total = _.map(this.stats.voters, "votes");
      let max = _.max(total);
      console.log(max);
      return max;
    }
  },
  methods: {
    ...mapActions(["fetchStats", "fetchSpeakers"]),
    percent(value) {
      return value;
    }
  },
  mounted() {
    this.fetchStats();
  }
};
</script>


<style lang="scss">
.page {
  h1,
  h2 {
    color: white;
  }

  h2 {
    font-size: 30px;
  }

  .speaker {
    width: 100%;
    max-width: 900px;
    display: grid;
    grid-template-columns: 2fr 3fr 1fr;
    grid-template-areas: "👶 📊 💯";
    margin-bottom: 5px;
  }

  .bar-label {
    font-size: 20px;
    text-align: left;
    padding-left: 10px;
  }

  .stats-wrapper {
    margin-bottom: 100px;
  }

  .name {
    grid-area: 👶;
  }

  .graph {
    grid-area: 📊;
  }

  .photo {
    width: 50px;
    height: 50px;
    border-radius: 30px;
  }

  .rectangle {
    //   width: 100%;
    .bar {
      background: white;
      height: 100%;
      content: " ";
      display: block;
      justify-content: flex-start;
      align-items: center;
      display: flex;
    }
  }

  .value {
    grid-area: 💯;
    color: white;
  }
}
</style>
