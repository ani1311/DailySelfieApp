document.getElementById('upload-form').addEventListener('submit', function (event) {
    event.preventDefault();

    var files = document.getElementById('image-upload').files;
    var imageDisplay = document.getElementById('image-display');

    for (var i = 0; i < files.length; i++) {
        var file = files[i];
        var img = document.createElement('img');
        img.src = URL.createObjectURL(file);
        imageDisplay.appendChild(img);

        // Create a FormData instance
        var formData = new FormData();
        // Append the file to the form data
        formData.append('image', file);

        // Send a POST request with the form data
        fetch('/upload', {
            method: 'POST',
            body: formData
        }).then(response => response.json())
            .then(data => console.log(data))
            .catch(error => console.error(error));
    }
});
