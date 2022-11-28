// You can read about that reactive implementation from:
// https://dev.to/themarcba/create-your-own-vue-js-from-scratch-part-3-building-the-reactivity-5162

let activeEffect;

// Adds tracking functionality to object properties
function reactive(obj) {
  Object.keys(obj).forEach((key) => {
    let dep = new Dep();
    let objValue = obj[key];

    Object.defineProperty(obj, key, {
      get() {
        dep.depend();
        return objValue;
      },

      set(newValue) {
        objValue = newValue;
        dep.notify();
      },
    });
  });
  return obj;
}

class Dep {
  // Initialize the value of the reactive dependency
  constructor() {
    this.subscribers = new Set();
  }

  // Subscribe a new function as observer to the dependency
  depend() {
    if (activeEffect) {
      this.subscribers.add(activeEffect);
    }
  }

  // Notify subscribers of a value change
  notify() {
    this.subscribers.forEach((subscriber) => subscriber());
  }
}

function watchEffect(fn, val) {
  activeEffect = fn;
  fn();

  // activeEffect = null;
}

export { watchEffect, Dep, reactive };
