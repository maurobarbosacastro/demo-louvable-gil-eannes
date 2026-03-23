import CryptoJS from 'crypto-js';

export function hashPassword(password) {
    return CryptoJS.AES.encrypt(password, '').toString();
}


// Function to validate password strength
export function createPasswordStrengthValidator(value) {
    if (!value) {
        return null; // returning a resolved promise
    }

    const hasUpperCase = /[A-Z]+/.test(value);
    const hasLowerCase = /[a-z]+/.test(value);
    const hasNumeric = /[0-9]+/.test(value);
    const hasSize = value.length >= 8;
    const hasSpecialChar = /[!@#$%^&*()_+\-=\[\]{};':"\\|,.<>\/?]/.test(value);

    return hasUpperCase && hasLowerCase && hasNumeric && hasSize && hasSpecialChar;
}
