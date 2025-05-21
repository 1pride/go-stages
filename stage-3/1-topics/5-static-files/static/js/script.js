document.addEventListener("DOMContentLoaded", function () {
    const form = document.getElementById("feedback-form");
    const status = document.getElementById("status");

    form.addEventListener("submit", async function (e) {
        e.preventDefault(); // Prevent page reload

        const formData = new FormData(form);
        const data = Object.fromEntries(formData.entries());

        try {
            const response = await fetch("/submit", {
                method: "POST",
                headers: {"Content-Type": "application/json"},
                body: JSON.stringify(data)
            });

            if (response.ok) {
                status.textContent = "Thanks for your feedback!";
                form.reset();
            } else {
                status.textContent = "Submission failed. Try again.";
            }
        } catch (err) {
            console.error(err);
            status.textContent = "An error occurred.";
        }
    });
});
