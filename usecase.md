# 🐦 Simple Twitter Clone — Use Case Specification

## 1. Introduction

### Purpose
This document describes the main **use cases** for the early-stage Twitter-like social media application.  
The goal is to define user interactions and system behaviors clearly before implementation.

### Scope
The system allows users to:
- Register and log in
- Post short text updates (up to 140 characters) with an optional image
- Follow and unfollow other users
- View a real-time feed of posts from followed users
- Receive notifications when followed users post new tweets
- Enjoy automatic live updates through **Server-Sent Events (SSE)**
- Include basic monitoring with **Prometheus** and **Grafana**

### Actors

| Actor | Description |
|--------|-------------|
| **User** | A registered individual who interacts with the system (posts, follows, views feed, etc.) |
| **System** | The backend services handling requests, data storage, and notifications |
| **Feed Service (SSE)** | Handles real-time updates and notifications |
| **Admin (optional)** | Used for monitoring and maintenance via Prometheus/Grafana |

### Assumptions
- Users must be logged in to interact with core features (post, follow, view feed).
- Pagination is not implemented in this version; feed is continuously loaded via scrolling.
- Notifications and live feed updates are delivered using Server-Sent Events (SSE).
- The system runs in Docker containers and exposes Prometheus metrics for monitoring.

---

## 2. Use Case List

| ID | Use Case | Description |
|----|-----------|-------------|
| UC-01 | Register / Login | User creates an account or logs in |
| UC-02 | Post Tweet | User creates a tweet (≤ 140 chars + optional image) |
| UC-03 | Follow User | User follows another user |
| UC-04 | Unfollow User | User stops following someone |
| UC-05 | View Feed | User views a live feed of tweets from followed users |
| UC-06 | Receive Notification | User receives a notification when a followed user posts |
| UC-07 | Real-time Feed Update (SSE) | System pushes new tweets to users’ feeds automatically |
| UC-08 | Logout | User ends their session |
| UC-09 | Monitoring & Metrics | Admin monitors system health with Prometheus & Grafana |

---

## 3. Detailed Use Cases

---

### **UC-01 — Register / Login**

**Primary Actor:** User  
**Goal:** Allow a user to create an account or log in to the system.

**Preconditions:**
- The user is not currently logged in.

**Postconditions:**
- User is authenticated and redirected to their feed.

**Main Flow:**
1. User opens the app.
2. Selects “Sign Up” or “Login”.
3. Provides credentials.
4. System validates input.
5. System logs the user in and opens the feed.

**Alternative Flows:**
- Invalid credentials → show error.
- Email or username already exists → show warning.

---

### **UC-02 — Post Tweet**

**Primary Actor:** User  
**Goal:** Post a tweet with up to 140 characters and an optional image.

**Preconditions:**
- User is logged in.

**Postconditions:**
- Tweet is stored in the system.
- Followers see the tweet in their feed in real time.

**Main Flow:**
1. User types tweet text (≤ 140 chars).
2. (Optional) attaches an image.
3. Clicks “Post”.
4. System validates content.
5. System stores tweet in the database.
6. System triggers SSE updates to followers.

**Alternative Flows:**
- Text too long → show error.
- Image upload fails → post without image.

---

### **UC-03 — Follow User**

**Primary Actor:** User  
**Goal:** Follow another user to receive their future tweets.

**Preconditions:**
- User must be logged in.
- Target user exists.

**Postconditions:**
- Follow relationship stored in database.
- Future tweets from the followed user will appear in feed.
- Notifications are enabled for this relationship.

**Main Flow:**
1. User visits another profile.
2. Clicks “Follow”.
3. System saves the relationship.
4. System confirms the follow action.

**Alternative Flows:**
- Already following → show “Already following” message.

---

### **UC-04 — Unfollow User**

**Primary Actor:** User  
**Goal:** Stop following a user.

**Preconditions:**
- User must be logged in.
- Follow relationship exists.

**Postconditions:**
- Follow relationship is removed.
- The unfollowed user’s tweets no longer appear in the feed.

**Main Flow:**
1. User visits followed user’s profile.
2. Clicks “Unfollow”.
3. System removes the relationship.
4. Confirmation is displayed.

---

### **UC-05 — View Feed**

**Primary Actor:** User  
**Goal:** View a real-time list of tweets from followed users.

**Preconditions:**
- User must be logged in.
- At least one tweet exists in the system.

**Postconditions:**
- User sees tweets from followed users in reverse chronological order.
- Feed updates automatically.

**Main Flow:**
1. User opens home page.
2. System retrieves tweets from followed users.
3. Feed Service (SSE) listens for new tweets.
4. Feed updates automatically when new tweets arrive.

---

### **UC-06 — Receive Notification**

**Primary Actor:** System  
**Goal:** Notify users when a followed user posts a tweet.

**Preconditions:**
- A follow relationship exists between users.

**Postconditions:**
- Notification delivered via SSE.

**Main Flow:**
1. User A posts a tweet.
2. System finds all followers of User A.
3. System sends notification to each follower via SSE.
4. Notification displayed in followers’ UI.

---

### **UC-07 — Real-time Feed Update (SSE)**

**Primary Actor:** Feed Service (SSE)  
**Goal:** Automatically push new tweets to users’ feeds in real time.

**Preconditions:**
- SSE connection is established.

**Postconditions:**
- User’s feed updates instantly when new tweets arrive.

**Main Flow:**
1. User opens the feed page.
2. SSE connection is opened.
3. When any followed user posts, backend triggers an SSE event.
4. Client updates the feed dynamically.

---

### **UC-08 — Logout**

**Primary Actor:** User  
**Goal:** End user session safely.

**Preconditions:**
- User is logged in.

**Postconditions:**
- User session is invalidated.
- Redirected to login page.

**Main Flow:**
1. User clicks “Logout”.
2. System clears session/token.
3. User redirected to login screen.

---

### **UC-09 — Monitoring & Metrics**

**Primary Actor:** Admin / DevOps  
**Goal:** Monitor the health and performance of the system.

**Preconditions:**
- Application is running in Dockerized environment.

**Postconditions:**
- Prometheus collects metrics.
- Grafana displays dashboards for system performance.

**Main Flow:**
1. System exposes Prometheus metrics endpoints.
2. Prometheus scrapes metrics.
3. Grafana visualizes real-time metrics.
4. Admin monitors CPU, memory, requests, and errors.

---

## 4. Future Use Cases (for later phases)

- UC-10 — Like a Tweet
- UC-11 — Comment / Reply
- UC-12 — Retweet
- UC-13 — Private Messaging
- UC-14 — Search & Hashtags

---

### ✅ Summary

This document defines the core interactions for the MVP version of the Twitter-like application.  
It is suitable for early implementation and future expansion, with SSE-based real-time behavior and observability in place.

---

© 2025 — Early Twitter Clone Project
