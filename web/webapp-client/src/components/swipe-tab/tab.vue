<script>
    export default {
        props: {
            label: {
                required: true
            },
            hash: {
                required: true
            },
            fontsize: {
                type: String
            },
            tabheight: {
                type: String
            },
            frontIcon: {
                type: String,
                default: ""
            },
            postIcon: {
                type: String,
                default: ""
            },
            RedDot: {
                type: String,
                default: ""
            },
            color: {
                type: String
            }
        },
        data() {
            return {
                isActive: false
            };
        },
        computed: {
            tabHeader() {
                return this.frontIcon + this.label + this.postIcon + this.RedDot;
            },
            tabHash() {
                return `${this.hash.toLowerCase()}`;
            },
            fontSize() {
                return parseInt(this.fontsize) / 75 + "rem";
            },
            tabHeight() {
                return parseInt(this.tabheight) / 75 + "rem";
            }
        }
    };
</script>

<template>
    <div class="tabs-panel-content">
        <slot />
    </div>
</template>

<style lang="scss">
    .tabs-panel-content {
        background: var(--color-white);
        padding: 0 10px;
        font-family: var(--font-glacial);

        .session-panes {
            font-size: 14px;
            list-style: none;
            margin: 0;
            border-bottom: 1px solid lightgrey;
            text-align: left;

            a {
                display: block;
                padding: 8px 0;
                text-decoration: none;
                color: white;
                transition: transform 0.2s;

                &:hover {
                    transition: transform 0.2s;
                    transform: translateX(5px);
                }
            }

            .date-time {
                text-transform: uppercase;
                font-size: 13px;
                font-family: var(--font-shentox);
                font-weight: 700;
                color: var(--color-green);
            }

            .session-title {
                padding: 5px 0;
                font-size: 20px;
                color: var(--color-blue);
                font-weight: 500;
                text-transform: uppercase;
            }

            .session-author {
                font-size: 14px;
                color: #333;
                text-transform: uppercase;
                font-weight: 300;
            }
        }
    }
</style>
