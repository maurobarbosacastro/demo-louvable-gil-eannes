importScripts(
    "https://www.gstatic.com/firebasejs/11.4.0/firebase-app-compat.js"
);
importScripts(
    "https://www.gstatic.com/firebasejs/11.4.0/firebase-messaging-compat.js"
);

const firebaseConfig = {
    apiKey: "AIzaSyBryVQZZ8MsnvGGhq96XMlgEafvQ7wuJ4U",
    authDomain: "tagpeak-test.firebaseapp.com",
    projectId: "tagpeak-test",
    storageBucket: "tagpeak-test.firebasestorage.app",
    messagingSenderId: "509240225527",
    appId: "1:509240225527:web:e015a320db03bde4466df2",
    measurementId: "G-ML325X3LFV"
};

const app = firebase.initializeApp(firebaseConfig);
const messaging = firebase.messaging();
