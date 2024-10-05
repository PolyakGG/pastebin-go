function validateTitle() {
    const maxLength = 30;
    const titleInput = document.getElementById('title');
    const charCountDiv = document.getElementById('charCount');

    const remainingChars = maxLength - titleInput.value.length;
    charCountDiv.textContent = "Осталось символов: " + remainingChars;
}