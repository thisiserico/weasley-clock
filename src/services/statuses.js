const address = "/.netlify/functions/statuses"

export const fetchEverything = () => {
  return fetch(address)
    .then(response => response.json())
}
