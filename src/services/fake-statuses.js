const statuses = [
  "dentist",
  "prison",
  "lost",
  "quidditch",
  "dying",
  "tailor",
  "bed",
  "holidays",
  "forest",
  "work",
  "garden",
  "school",
  "home",
]

const randomStatus = () => {
  const index = Math.floor(Math.random() * Math.floor(statuses.length))
  return statuses[index]
}

export const fetchEverything = () => {
  return new Promise(resolve => {
    resolve({
      statuses: statuses,
      people: {
        ron: {
          name: "Ron",
          status: randomStatus()
        },
        ginny: {
          name: "Ginny",
          status: randomStatus()
        },
        george: {
          name: "George",
          status: randomStatus()
        },
        fred: {
          name: "Fred",
          status: randomStatus()
        }
      }
    })
  })
}
