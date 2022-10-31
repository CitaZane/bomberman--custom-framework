
export default {
    state: {
        userQueueCount: 0,
        messages: []
    },
    mutations: {
        updateUserQueueCount(state, count) {
            // console.log("UPDATED USER COUNT")
            state.userQueueCount = count
        }
    },
    actions: {
        // addNewMessage({state, commit}, message) {
        //     const messages = state.messages;
        //     messages.push(message);

        // }
    },
}

