const API_BASE_URL = "http://localhost:8080/api";
let incomes = [];
let expenses = [];
let userID;

function getHeaders() {
  return {
    "Content-Type": "application/json",
    Accept: "application/json",
    Authorization: localStorage.getItem("authHeader"),
  };
}

function formatDate(dateString) {
  const d = new Date(dateString);
  const day = String(d.getDate()).padStart(2, "0");
  const month = String(d.getMonth() + 1).padStart(2, "0");
  const year = d.getFullYear();
  return `${day}/${month}/${year}`;
}

async function loadData(type, targetArray) {
  try {
    const response = await fetch(`${API_BASE_URL}/${type}`, {
      method: "GET",
      headers: getHeaders(),
    });

    if (!response.ok) {
      throw new Error("Unable to load list");
    }

    const data = await response.json();
    targetArray.length = 0;
    if (type === "income") {
      data.data.incomes.forEach((item) => targetArray.push(item));
    } else {
      data.data.expenses.forEach((item) => targetArray.push(item));
    }
  } catch (error) {
    alert("Something went wrong");
    console.error(error);
  }
}

async function renderData() {
  await loadData("income", incomes);
  await loadData("expense", expenses);

  // Sau khi load xong, render ra bảng
  incomes.forEach((income) => {
    const data = {
      date: formatDate(income.date),
      description: income.description,
      amount: income.amount,
    };
    console.log(data.date);
    createRow(document.querySelector("#income-table tbody"), data, "income");
  });

  expenses.forEach((expense) => {
    const data = {
      date: formatDate(expense.date),
      description: expense.description,
      amount: expense.amount,
    };
    console.log(data.date);
    createRow(document.querySelector("#expense-table tbody"), data, "expense");
  });

  // Cập nhật tổng
  totalIncome = incomes.reduce((sum, i) => sum + i.amount, 0);
  totalExpense = expenses.reduce((sum, e) => sum + e.amount, 0);
  updateTotals();
}

document.addEventListener("DOMContentLoaded", function () {
  if (!localStorage.getItem("authHeader")) {
    window.location.href = "index.html";
    return;
  }

  fetch(`${API_BASE_URL}/profiles`, {
    method: "GET",
    headers: getHeaders(),
  })
    .then((res) => res.json())
    .then((response) => {
      userID = response.data.user.id;
      const fullName = response.data.user.full_name;
      const username = response.data.user.username;
      document.getElementById("fullname").textContent = fullName;
      document.getElementById("username").textContent = username;
    })
    .catch((error) => {
      console.error("Lỗi:", error);
    });

  renderData();
});

async function submitTransaction(e, type) {
  e.preventDefault();
  const form = e.target;

  const data = {
    date: form.date.value,
    description: form.description.value,
    amount: parseInt(form.amount.value),
    user_id: userID,
  };

  console.log(`[${type.toUpperCase()}] Gửi dữ liệu:`, data);

  try {
    const res = await fetch(`${API_BASE_URL}/${type}`, {
      method: "POST",
      headers: getHeaders(),
      body: JSON.stringify(data),
    });

    if (!res.ok) throw new Error(`Lỗi khi gửi ${type} lên server`);

    const response = await res.json();
    console.log(`[${type.toUpperCase()}] Đã lưu vào DB:`, response);

    createRow(document.querySelector(`#${type}-table tbody`), data, type);
    updateTotals();
    form.reset();
  } catch (error) {
    console.error(`Lỗi khi thêm ${type}:`, error);
    alert(`Thêm ${type === "income" ? "thu nhập" : "chi tiêu"} thất bại!`);
  }
}

document
  .getElementById("income-form")
  .addEventListener("submit", (e) => submitTransaction(e, "income"));

document
  .getElementById("expense-form")
  .addEventListener("submit", (e) => submitTransaction(e, "expense"));
