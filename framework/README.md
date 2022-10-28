<!-- If you are looking at documentation in raw format, press ctrl+shift+v to see preview -->

# Mini-Framework

- [Mini-Framework](#mini-framework)
  - [Introduction](#introduction)
  - [Quick Start](#quick-start)
  - [Creating an Application](#creating-an-application)
    - [The Root Component](#the-root-component)
  - [Reactivity](#reactivity)
  - [Templating / JSX](#templating--jsx)
    - [Embedding Expressions](#embedding-expressions)
    - [Specify Attributes](#specify-attributes)
  - [Components](#components)
    - [Defining a component](#defining-a-component)
    - [Using a component](#using-a-component)
  - [Component lifecycle](#component-lifecycle)
  - [Refs](#refs)
  - [Handling Events](#handling-events)
  - [Conditional Rendering](#conditional-rendering)
  - [List Rendering](#list-rendering)
  - [State](#state)
    - [Persisting state](#persisting-state)
  - [Routing](#routing)
    - [Dynamic Route Matching](#dynamic-route-matching)

## Introduction

Mini-framework is a JavaScript framework for making it easy to build interactive user interfaces.
Though built as an higher abstraction level, still basic familiarity with HTML, CSS and JavaScript is needed.
Frameworks core features are

- Reactivity: tracking JavaScript state changes and efficiently updating the DOM by using Virtual DOM
- Declarative Rendering: by using [JSX](#template--jsx) it is possible to change HTML output based on JavaScript state
- Component based architecture: allowing bundling together HTML with JavaScript logic in reusable blocks of code
- Integrated Routing System
- State Management

## Quick Start

At the root of the project("/mini-framework") run those commands:

```
npm install
```

This will install needed packages to run the application

After installation start a dev server with

```
npm run dev
```
After server has successfuly started [open application](http://localhost:1234/) in browser.

## Creating an Application
Your appliction is located in `/src` folder.
Every application starts by initializing router and store:

```js
//app.js

// import user defined routes and storeObject
import storeObj from "./store/index";
import routes from "./router/index";
// import initialization functions
import createStore from "../framework/store";
import createRouter from "../framework/router";

const store = createStore(storeObj);
const router = createRouter(routes);

export { store, router };
```

### The Root Component

Every app requires a "root component" that can contain other components as its children. As framework comes with inbuilt router the root component is imported in router and passed in first defined route. Root componets are the main enterypoint for application, though they internally work the same as other components, for application organizational purposes they are separated and defined as views. Define view in views directory:

```jsx
// views/HomeView.js
import jsx from "../../framework/vDom/jsx";

export const HomeView = () => {
  return {
    template: <h1> This is home view</h1>
  }
};
```

Import defined view in router and pass as root component for route:

```js
// router/index.js
import { HomeView } from "../views/HomeView";

const routes = [{ path: "/", component: HomeView }];

export default routes;
```

Root component on first render will be mounted inside predefined <div></div> element with id = "root". This element with id should be present in html file for framework to work correctly:

```html
<!-- index.html -->

<body>
  <div id="root"></div>
  <script type="module" src="app.js"></script>
</body>
```

In an simple application there can be only one single component or single view, but usually applications are larger and organized into nested trees of reusable components. For example, a Todo application's tree might look like this:

```
todo-app
├─ store
│  └─ index.js
├─ router
│  └─ index.js
├─ css
│  └─ main.css
├─ views
│  ├─ HomeView.html
│  ├─ ActiveView.html
│  └─ CompletedView.html
├─ components
│  ├─ TodoList.js
│  └─ TodoItem
│       ├─ TodoItem.js
│       ├─ TodoDeleteButton.js
│       └─ TodoEditButton.js
├─ app.js
└─ index.html
```

## Reactivity

Our reactivity system is inspired by Vue's. We make an object reactive by calling a reactive function and passing it the object we want to make reactive.

```js
import { watchEffect, reactive } from "../../framework/reactive";

// create a reactive object
const reactiveObj = reactive({
  name: "John",
  age: 50,
});
```

To make something happen when our object values change. We have to call **watchEffect**.<br>
**watchEffect** accepts a function and it will set the given function as an **activeEffect**.

```js
watchEffect(() => {
  console.log(reactiveObj.name, reactiveObj.age);
});
```

Every time we read a value from an object an **activeEffect** will be saved for that values key.

Now if we set a new value for an object key the previously set **activeEffect** will be called.

```js
reactiveObj.name = "Mike";
reactiveObj.age = 10;
// output:
// Mike, 50
// Mike, 10
```
Under the hood this reactivity system is used to track the values in user defined store.state and update the DOM whenever chenges are detected.
## Templating / JSX

To describe how UI should look like framework uses JSX. While from the first glance it can be seen as a template language, it is much stronger as it takes full advantage of JavaScript. A simple variable declaration looks like this:

```jsx
const element = <h1>Hello, world!</h1>;
```

JSX opens door to bulding simple components by coupling logic together with markdown. Components are described in detail in next section, but lets take a look at some useful JSX features:

### Embedding Expressions

You can declare a variable and then use it inside JSX by wrapping it in curly braces:

```jsx
const answer = 42;
const element = <h1>The ultimate answer is: {answer}</h1>;
```

We can go even further than variables, you can put any valid JavaScript expression inside the curly braces. From simple arithemetic or logical expression to calling a function.

### Specify Attributes

Using curly braces JavaScript expressions can be embeded in an attribute. One more thing to keep in mind, JSX uses `camelCase` property naming convention instead of HTML attribute names:

```jsx
const element = <img src={user.avatar} className="avatar"></img>;
```

## Components

Components allow us to split the UI into independent and reusable pieces, and think about each piece in isolation. It's common for an app to be organized into a tree of nested components.

### Defining a component

Conceptually, components are like JavaScript functions. They accept inputs and return JSX elements as a template. In each component file `jsx` should be imported for component to work correctly:

```jsx
// components/SimpleComponent.js
import jsx from "../../framework/vDom/jsx";

export function SimpleComponent(props) {
  return {
    template:(
      <h1>Hello, {props.name}</h1>
    )
  };
}
```

Elements can represent not only DOM tags, but also user-defined components:

```jsx
const element = <Welcome name="John" />;
```

### Using a component

To use user-defined components they should be brought into scope(imported).
By using custom attributes you can pass properties to child components:

```jsx
import jsx from "../../framework/vDom/jsx";
import { TodoList } from "./TodoList";

export const TodoApp = () => {
  const todoData = [
    { task: "one", id: 1 },
    { task: "two", id: 2 },
    { task: "three", id: 3 },
    { task: "four", id: 4 },
  ];
  return {
    template: (
      <h1>todos</h1>
      <TodoList list={todoData} />
    )
  }
};
```

After passing data as props to child component, you can destructure by attribute name, and use the data:

```jsx
export const TodoList = ({list}) => {
  return {
    template: //you can use list data to render neccassary elements
  }
};
```

## Component lifecycle

You have the option to run a function when the component is attached to real DOM.

To do that you have to add an additional property **onMounted** with a value of function which will then be called.

```js
export function TodoItem({ toDo }) {
  return {
    onMounted: () => {
      console.log("TodoItem component mounted!");
    },
    template: (
      <li class={todoItemClass(toDo)} data-id={toDo.id}>
        <div class="view">/* some jsx */</div>
      </li>
    ),
  };
}
```

## Refs

Refs are references to real HTML elements. You can use refs to interact with DOM.

For example you can save a reference of an input element and focus on it once it is rendered.

Creating a reference to an element is easy, you just have to add an attribute called **ref** and give it a string name, which you will later use to access the element.

```html
<input type="text" ref="editText" />
```

You can access refs only in the onMounted lifecycle, which will be provided to the lifecycle function as an object by the framework.

```js
export function EditItem({ toDo }) {
  return {
    template: (
      <input type="text" class="edit" value={toDo.task} ref="editText" />
    ),

    onMounted: (refs) => {
      console.log("Edit mode component mounted!");

      // access the input HTML element
      const textInput = refs.editText;

      // change the element state
      textInput.focus();
      textInput.setSelectionRange(
        textInput.value.length,
        textInput.value.length
      );
    },
  };
}
```

## Handling Events

For event handling keep in mind two main priciples:

- Use `camelCase` syntax for event names
- Pass a function as event handler

After applying these two principles, the event handling is pretty straightforward:

```jsx
export const AddItem = () => {
  function handleKeyup(event) {
    if (event.keyCode === 13) {
      alert("Adding item...");
    }
  }
  return {
    template: (
      <input onKeyup={handleKeyup}></input>
    )
  }
};
```

To pass arguments to event handler use an arrow function:

```jsx
function deleteTodoItem(e, toDo) {
  //do the destroying
}

export const DestroyItem = ({ toDo }) => {
  return {
    template: (
      <li>
        <button onClick={(e) => deleteTodoItem(e, toDo)}></button>
      </li>
    )
  };
};
```

## Conditional Rendering

Conditional rendering can be done in three ways, based on your needs.
You can render blocks of code or components based on JavaScript logic:

```jsx
function Greeting(props) {
  const isLoggedIn = props.isLoggedIn;
  if (isLoggedIn) {
    return {template: <UserGreeting />};
  }
  return {template: <GuestGreeting />};
}
```

You can use inline logical operator `&&` to conditionaly render element,by wrapping the expression, operator and element in curly braces:

```jsx
function Greeting(props) {
  const isLoggedIn = props.isLoggedIn;
  return {
    template: (
      <div>
        <h1>Hello!</h1>
        {isLoggedIn && <h2>{props.username}</h2>}
      </div>
    );
  }
}
```

And to implement `if-else` type rendering, use `condition ? true : false`

```jsx
function Greeting(props) {
  const isLoggedIn = props.isLoggedIn;
  return{
    template: (
      {isLoggedIn ? <UserGreeting /> :  <GuestGreeting />}
    )
  }
}
```

## List Rendering

You can use list rendering to create multiple elements based on array or object.
Lets take to-do application as an example. It would be necassary to render each to-do from a list, and it could be done using `map()` function:

```jsx
export const TodoList = () => {
  const todoData = [
    { task: "one", id: 1 },
    { task: "two", id: 2 },
    { task: "three", id: 3 },
    { task: "four", id: 4 },
  ];

  const list = toDoData.map((item) => <TodoItem toDo={item} />);

  return {template: <ul>{list}</ul>};
};
```

```jsx
export function TodoItem({ toDo }) {
  return {
    template: (
      <li id={toDo.id}>
        <p>{toDo.task}</p>
      </li>
    )
  };
}
```

Similar to list rendering based on arrays, you can also use objects. In this case simple `map()` function is no longer valid, but it can be used if you turn object into array based on your needs, by using `Object.entries()`if both key and values are needed, or `Object.keys()` if you need to pass data about keys, or `Object.values()`, if only values needed:

```jsx
let list = Object.entries(someObject).map((item) => (
  <ChildComponent item={item} />
));
```

## State

State in our framework is managed globally via **Store**, which consists of:

- **State**     -> Holds the data
- **Mutations** -> Updates/mutates the state.
- **Actions**   -> Change state by calling mutations.
  
```js
// store/index.js

// define store
const storeObj = {
  state: {
    todoList: [{ task: "one", id: 1, completed: false }],
  },
  mutations: {
    updateTodoList(state, todoList) {
      state.todoList = todoList;
    },
  },
  actions: {
    addTodoItem({ state, commit }, todoItem) {
      let todoList = state.todoList;
      todoList.push(todoItem);
      // calling a mutation
      commit("updateTodoList", todoList);
    },
  },
};
export default storeObj;
```

You can initialize the Store in app.js by calling **createStore**, which will accept an object, which will have:

1. State
   - Object with data
2. Mutations
   - Object with functions
   - Each function will accept two arguments: state and data
     - State is the stores state object, which will be provided by the framework
     - Data is the data you want to update the store with and will be provided by the caller
3. Actions
   - Object with functions
   - Each function will accept two arguments: store object and data
     - State is the stores state object, which will be provided by the framework
     - Store object is provided by the framework and will hold the store's state and a function to commit mutations
```js
// app.js

import storeObj from "./store/index";
import createStore from "../framework/store";

const store = createStore(storeObj);

export { store };
```

To invoke an action, we call dispatch with action name and data.

```js
dispatch("actionName", data);
```

To invoke a mutation, we call commit with mutation name and data.

```js
commit("mutationName", data);
```
Example with invoking an action from component:
```js
// components/AddItem.js

import jsx from "../../framework/vDom/jsx";
import { store } from "../app";

export const AddItem = () => {
  function addItemHandler(event) {
     // .. //

    const todoItem = {
        task: event.target.value,
        id: Date.now(),
        completed: false,
    };

    // call an action addTodoItem with todoItem from the store
    store.dispatch("addTodoItem", todoItem);

    // ... //
    }
    return {
      template: (
        <input
          onKeyup={(e) => addItemHandler(e)}
          class="new-todo"
          placeholder="What needs to be done?"
          autofocus=""
        ></input>
      )
    };
  }
```
### Persisting state
To give user more controll, state persisting is not inbuilt in framework, but it can be easly added by using localStorage and including something similar to this in app.js:
```js

// save store.state.<some-data> in localStorage before unload
window.onbeforeunload = () => {
  localStorage.setItem("framework-data", JSON.stringify(store.state.data));
};

// access persisted data and save in store
window.onload = () => {
  let persistedData = JSON.parse(localStorage.getItem("framework-data"));

  if (persistedData) {
    store.state.data = persistedData.data;
  }
};
```

## Routing

Router is implemented in mini-frameworks core. It is used to map components(views) to routes and let router know when to render them. All routes should be defined in router/index.js file.
First defined route will be set as a default route, in case of no match will be found, view that is defined on fist route will be rendered. To keep things organized, the components that are defined as root elements in router should be called `views` and defined in `views` directory.

```js

// Import view components
import { HomeView } from "../views/HomeView";
import { AboutView } from "../views/AboutView";

// define routes and  initialize the router
const routes = [
  { path: "/", component: HomeView },
  { path: "/about", component: AboutView },
];
// export the routes
// to be imported in app.js
export default routes;
```

Initialize router in app.js :
```js
import routes from "./router/index";
import createRouter from "../framework/router";

const router = createRouter(routes);

export { router };
```
By importing router in component we can access current route as `router.currentRoute`.

```jsx
// HomeView.js
import { router } from "../app";

export const HomeView = () => {
  // access properties from current route
  // for more about properties look into Dynamic Route Matching
  let username = router.currentRoute.props.username;

  return (
    <section>
      <h1>Welcome back, {username}</h1>
    </section>
  );
};
```

To create links use regular anchor tags, but in front of reference link include `#`:

```js
<a href="#/about">About</a>
```

### Dynamic Route Matching

To match routes with given pattern to the same component you can use dynamic segments. For example to render `User` component, to different users, who can have different id's inside of the path. So urls like `/user/Jane` and `/user/Peter` would both map to the same route, while saving the dynamic part to be accessed later inside the component.

```js
// router/index.js

// dynamic segment starts with a colon
{ path:"/user/:username", component: UserView }
```

```js
// views/UserView.js
import { router } from "../app";

export const UserView = () => {
  // access properties from current route
  let username = rouer.currentRoute.props.username;

  return {
    template: (
      <section>
        <h1>And the current user is: {username}</h1>
      </section>
    )
  };
};
```

One route can hold multiple dynamic segments, as long as their names do not match.

```js
/// router/index.js
{ path:"/user/:username/posts/:postId", component: PostView }
```
