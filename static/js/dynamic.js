// Update the hx-get attribute of the select element
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