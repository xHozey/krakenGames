const container = document.getElementById('container')

const getData = async () => {
    try {
        let response = await fetch("http://localhost:8080/api")
        let data = await response.json()
        displayPosts(data)
    } catch (err) {
        console.error(err)
    }
}

const displayPosts = (data) => {
    data.forEach(element => {
        let card = document.createElement('div')
        card.innerHTML = `
        <span class="title">${element.title}</span>
        <a href="${element.id}"><img class="card-image" src="${element.image.split('/assets').join('')}"/></a>
        `
        card.className = "boarder"
        container.append(card)
    });
}

getData()