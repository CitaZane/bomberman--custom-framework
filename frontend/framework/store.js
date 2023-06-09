/* 
    Link for reference: https://medium.com/@lachlanmiller_52885/building-vuex-from-scratch-9ac47c768f9d
*/
import { reactive } from "./reactive";
class Store {
  constructor({ state, mutations, actions }) {
    this.state = reactive(state);
    this.mutations = mutations;
    this.actions = actions;
  }

  commit(handler, payload) {
    this.mutations[handler](this.state, payload);
  }

  dispatch(handler, payload) {
    // bind "this" value to commit function because value of "this" depends where function is called
    return Promise.resolve(
      this.actions[handler](
        { state: this.state, commit: this.commit.bind(this) },
        payload
      )
    );
  }
}

export default function createStore(store) {
  return new Store(store);
}
