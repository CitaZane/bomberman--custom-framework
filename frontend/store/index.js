export default {
    state: {
        userQueueCount: 0,
        count: 0,
    },
    mutations: {
        updateUserQueueCount(state, count) {
            // console.log("UPDATED USER COUNT")
            state.userQueueCount = count
        }
    },
    actions: {},
}