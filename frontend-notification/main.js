function subscribe() {
    fetch("http://localhost:8080/api/notifications/subscribe", {
        method: "POST",
        headers: {
            "Authorization": localStorage.getItem("token") || "",
        },
    })
        .then(res => res.json())
        .then(() => showToast(" Подписка оформлена"))
        .catch(() => showToast(" Ошибка при подписке"));
}

function unsubscribe() {
    fetch("http://localhost:8080/api/notifications/unsubscribe", {
        method: "POST",
        headers: {
            "Authorization": localStorage.getItem("token") || "",
        },
    })
        .then(res => res.json())
        .then(() => showToast(" Вы отписались"))
        .catch(() => showToast(" Ошибка при отписке"));
}

function createNotification(event) {
    event.preventDefault();

    const title = document.getElementById("title").value;
    const message = document.getElementById("message").value;
    const street = document.getElementById("street").value;
    const sendAt = new Date(document.getElementById("send_at").value).getTime() / 1000;

    fetch("http://localhost:8080/api/notifications/create", {
        method: "POST",
        headers: {
            "Authorization": localStorage.getItem("token") || "",
            "Content-Type": "application/json",
        },
        body: JSON.stringify({ title, message, street, send_at: sendAt }),
    })
        .then(res => res.json())
        .then(() => showToast(" Уведомление создано"))
        .catch(() => showToast(" Ошибка при создании"));
}

function loadNotificationHistory() {
    fetch("http://localhost:8080/api/notifications/history", {
        headers: {
            "Authorization": localStorage.getItem("token") || "",
        },
    })
        .then(res => res.json())
        .then(data => {
            const container = document.getElementById("notification-history");
            container.innerHTML = data.map(n => `
        <p><strong>${n.title}</strong>: ${n.message}<br>
         ${new Date(n.send_at * 1000).toLocaleString()}<br>
         ${n.street}</p>
      `).join("");
        })
        .catch(() => {
            document.getElementById("notification-history").textContent = " Не удалось загрузить историю";
        });
}
