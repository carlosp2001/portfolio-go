document.addEventListener('DOMContentLoaded', () => {
    const nameInput = document.getElementById('name');
    const emailInput = document.getElementById('email');
    const subjectInput = document.getElementById('subject');
    const messageInput = document.getElementById('message');
    const submitButton = document.getElementById('submit');
    const successMessage = document.getElementById('success-message');

    submitButton.addEventListener('click', sendEmail)
    nameInput.addEventListener('input', () => {
        validateInput(nameInput);
    });

    emailInput.addEventListener('input', () => {
        validateInput(emailInput);
    })

    subjectInput.addEventListener('input', () => {
       validateInput(subjectInput);
    })

    messageInput.addEventListener('input', () => {
        validateInput(messageInput);
    })

    function sendEmail() {
        const name = nameInput.value;
        const email = emailInput.value;
        const subject = subjectInput.value;
        const message = messageInput.value;

        validateInput(nameInput);
        validateInput(emailInput);
        validateInput(subjectInput);
        validateInput(messageInput);

        const formData = {
            name,
            email,
            subject,
            message
        }

        fetch('/api/contact', {
            method: 'POST',
            body: JSON.stringify(formData),
            headers: {
                'Content-Type': 'application/json'
            }
        }).then(response => {
            if (response.ok) {
                successMessage.classList.remove('success-message--hidden');
                setTimeout(() => {
                    successMessage.classList.add('success-message--hidden');
                }, 3000);
            }
        }).catch(error => {
            console.log(error)
        })
    }
})

function validateInput(input) {
    if (input.value !== '') {
        input.classList.remove('label-default--error');
    } else {
        input.classList.add('label-default--error');
    }
}

