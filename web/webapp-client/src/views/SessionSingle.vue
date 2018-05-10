<template>
    <div class="page page-session">
        <div class="page-content" v-if="session">
            <!-- <span>{{id}}</span> -->
            <div class="session-title">{{session.title}}</div>

            <div class="speakers-wrapper" v-if="session.speakers">
                <div class="speaker-wrapper" v-for="speaker in session.speakers" :key="speaker.id">
                    <div class="avatar"><img :src="getSpeaker(speaker.id)" alt=""></div>
                    <p class="name">{{ speaker.name }}</p>
                </div>
            </div>

            <div class="descriptions-row">
                <div class="des-wrap" v-if="session.format">
                    <label>
                        <img src="../assets/icons/language.svg" alt="">
                    </label>
                    <p>{{ session.format }}</p>
                </div>

                <div class="des-wrap" v-if="session.language">
                    <label>
                        <img src="../assets/icons/language.svg" alt="">
                    </label>
                    <p>{{ session.language }}</p>
                </div>

                <div class="des-wrap">
                    <label>
                        <img src="../assets/icons/location.svg" alt="">
                    </label>
                    <p>{{ session.room }}</p>
                </div>

                <div class="des-wrap">
                    <label>
                        <img src="../assets/icons/time.svg" alt="">
                    </label>
                    <p>{{ time(session.startsAt) + getDay(session.startsAt) }} - {{ time(session.endsAt) }}</p>
                </div>

                <div class="des-wrap" v-if="session.level">
                    <label>
                        <img src="../assets/icons/level.svg" alt="">
                    </label>
                    <p>{{ session.level }}</p>
                </div>
            </div>

            <div class="description-text">
                <p>{{ session.description }}</p>
            </div>

            <div class="button-wrappers">
                <router-link :to="{ name: 'sessions' }" class="btn small back">
                    Back
                </router-link>

                <router-link :to="{ name: 'feedback',  params: { id: id }}" class="btn small rate">
                    Rate
                </router-link>
            </div>

        </div>
        <div class="page-content" v-else>
            calling Ish..
        </div>
        <div class="footer">
            Developer Conference 2018
        </div>

    </div>
</template>

<script>
    import { mapGetters, mapActions } from "vuex";
    import moment from "moment";

    export default {
        props: ["id"],
        methods: {
            ...mapActions(["fetchSessions", "fetchSpeakers"]),
            getSpeaker: function(id) {
                if (this.speakers.length === 0) {
                    this.fetchSpeakers();
                }
                console.log(id);
                console.log(this.speakers);
                let theSpeaker = this.speakers.filter(speaker => speaker.id === id);
                if (theSpeaker.length > 0) {
                    return theSpeaker[0].profilePicture;
                }
            },
            time: function(date) {
                // console.log()
                // return new Date(date).toDateString();
                return moment(date).format("LT");
            },
            getDay: function(str) {
                return str.split(",")[0];
            }
        },
        computed: {
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
        beforeMount() {
            if (this.$store.state.sessions.length === 0) {
                // console.error("no sessions found");
                this.fetchSessions();
            } else {
                // console.info("sessions found !");
            }
        }
    };
</script>

<style lang="scss" scoped>
    .page-session {
        display: grid;
        grid-template-areas: "session session" "footer footer";
        grid-template-columns: 100px 1fr;
        grid-template-rows: 1fr 10vh;
        max-width: 900px;
        margin: 0 auto;
        width: 100%;
        margin-top: 50px;
    }

    .page-content {
        grid-area: session;
        background: white;
        box-shadow: 0 0 10px rgba(0, 0, 0, 0.2);
    }

    .session-title {
        color: #333333;
        text-transform: uppercase;
        font-family: var(--font-glacial);
        font-size: 40px;
        font-weight: 700;
        margin: 0 auto;
        padding: 30px 10px;
        text-align: center;
    }

    .speakers-wrapper {
        display: flex;
        align-items: center;
        justify-content: center;
        margin-bottom: 20px;

        .speaker-wrapper {
            --width: 70px;
            display: grid;
            width: 200px;
            grid-template-areas: "avatar name";
            grid-template-columns: var(--width) 1fr;
            align-items: center;
            grid-gap: 10px;
            margin-bottom: 10px;
            margin-right: 20px;
            font-family: var(--font-shentox);
            font-weight: 500;
            text-transform: uppercase;

            .avatar {
                grid-area: avatar;
                width: var(--width);
                height: var(--width);
                border-radius: var(--width);
                box-shadow: 0 0 20px rgba(0,0,0,.1);
                overflow: hidden;

                img {
                    width: var(--width);
                    height: var(--width);
                    border-radius: var(--width);
                    display: block;
                    object-position: 50% 50%;
                    object-fit: cover;
                }
            }

            p {
                grid-area: name;
                text-align: left;
                font-size: 15px;
                color: var(--color-blue);
                font-weight: 700;
            }
        }

    }

    .descriptions-row {
        background: var(--color-blue);
        width: 100%;
        display: flex;
        align-items: center;
        font-size: 14px;
        padding: 10px;

        .des-wrap {
            min-width: 20%;
            margin-right: 10px;
            color: white;
            font-family: var(--font-glacial);
            text-transform: uppercase;
            display: flex;
            align-items: center;

            label {
                margin-right: 10px;

                img {
                    width: 20px;
                }
            }

            p {
                margin: 0;
            }
        }

    }

    .description-text {
        padding: 10px;

        p {
            font-family: var(--font-glacial);
            font-size: 18px;
            line-height: 25px;
            font-weight: 300;
        }
    }

    .button-wrappers {
        display: grid;
        grid-template-columns: 1fr 1fr;
        grid-template-rows: 1fr;

        .back {
            grid-column: 1 / 2;
        }

        .rate {
            grid-column: 2 / 3;
        }
    }

    .footer {
        grid-area: footer;
        color: black;
        text-transform: uppercase;
        font-family: var(--font-shentox);
        color: white;
        text-align: center;
        font-size: 13px;
        align-self: center;
    }
</style>

