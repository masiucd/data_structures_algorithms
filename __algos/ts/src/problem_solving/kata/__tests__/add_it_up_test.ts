import {assertEquals} from "../../../dev_deps.ts"
import {addItUp} from "../add_it_up/solution.ts"

const arr1 = [
  {name: "Joe Brown", goals: 0, assists: 0, points: 0},
  {name: "Jim Bob", goals: 2, assists: 1, points: 3},
  {name: "Harry Styles", goals: 1, assists: 1, points: 2},
  {name: "Craig Mack", goals: 5, assists: 7, points: 12},
  {name: "Marcell Ciszek Druzynski 🔥", goals: 5, assists: 7, points: 12},
  {name: "DOG", bones: 100000, goals: 5, assists: 7, points: 12},
]

const arr2 = [
  {
    name: "Craig Mack",
    goals: 3,
    assists: 3,
    points: 6,
    ppg: 0,
    ppa: 0,
    pims: 0,
  },
  {name: "Jim Bob", goals: 1, assists: 4, points: 5, ppg: 0, ppa: 1, pims: 0},
  {name: "Joe Brown", goals: 0, assists: 0, points: 0, ppg: 0, ppa: 0, pims: 0},
  {
    name: "Harry Styles",
    goals: 0,
    assists: 0,
    points: 0,
    ppg: 0,
    ppa: 0,
    pims: 0,
  },
]

Deno.test("addItUp it works!", () => {
  const result = addItUp(arr1, arr2)
  assertEquals(result, {
    "Joe Brown": {goals: 0, assists: 0, points: 0, ppg: 0, ppa: 0, pims: 0},
    "Jim Bob": {goals: 5, assists: 6, points: 11, ppg: 0, ppa: 1, pims: 0},
    "Harry Styles": {goals: 2, assists: 2, points: 4, ppg: 0, ppa: 0, pims: 0},
    "Craig Mack": {goals: 13, assists: 17, points: 30, ppg: 0, ppa: 0, pims: 0},
    "Marcell Ciszek Druzynski 🔥": {goals: 10, assists: 14, points: 24},
    DOG: {bones: 200000, goals: 10, assists: 14, points: 24},
  })
})
