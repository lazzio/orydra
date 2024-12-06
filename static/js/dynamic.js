// Update the hx-get attribute of the select element
// and load the client details
function updateHxGet(selectElement) {
    const selectedValue = selectElement.value;
    if (selectedValue) {
        selectElement.setAttribute('hx-get', `/api/client/${selectedValue}`);
        console.log(selectedValue);
        htmx.ajax('GET', `/api/client/${selectedValue}`, {
            target: '#client-details',
            swap: 'innerHTML'
        });
        document.getElementById('client-details').style.display = 'block';
    } else {
        selectElement.removeAttribute('hx-get');
        // redirect to the index page
        window.location.href = '/';
    }
}

function setupClientCreation() {
    const form = document.getElementById('createClientForm');
    if (!form) return;

    form.addEventListener('submit', function(e) {
        e.preventDefault();
        const formData = new FormData(form);
        const clientName = formData.get('client_name');
        
        // Désactiver le bouton pendant la création
        const submitButton = form.querySelector('button[type="submit"]');
        submitButton.disabled = true;
        
        // Créer une source d'événements
        const eventSource = new EventSource(`/api/client/create?client_name=${encodeURIComponent(clientName)}`);
        
        eventSource.onmessage = function(event) {
            const data = JSON.parse(event.data);
            
            if (data.error) {
                document.getElementById("message-container").innerHTML = 
                    `<div class="notification is-danger mt-6">${data.error}</div>`;
            } else {
                document.getElementById("client-id").textContent = data.clientId;
                document.getElementById("client-secret").textContent = data.clientSecret;
                document.getElementById("client-details").style.display = "block";
                document.getElementById("message-container").innerHTML = 
                    '<div class="notification is-success mt-6">Client created successfully!</div>';
            }
            
            // Réactiver le bouton
            submitButton.disabled = false;
            eventSource.close();
        };
        
        eventSource.onerror = function() {
            document.getElementById("message-container").innerHTML = 
                '<div class="notification is-danger mt-6">Error creating client</div>';
            // Réactiver le bouton
            submitButton.disabled = false;
            eventSource.close();
        };
    });
}

document.addEventListener('DOMContentLoaded', setupClientCreation);

