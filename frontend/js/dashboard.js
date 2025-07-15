let totalIncome = 0;
let totalExpense = 0;

function formatCurrency(amount) {
  return amount.toLocaleString("vi-VN") + "₫";
}

function updateTotals() {
  document.getElementById("total-income").textContent =
    formatCurrency(totalIncome);
  document.getElementById("total-expense").textContent =
    formatCurrency(totalExpense);
  document.getElementById("balance").textContent = formatCurrency(
    totalIncome - totalExpense
  );
}

function createRow(tableBody, data, type) {
  const row = document.createElement("tr");
  row.innerHTML = `
    <td>${data.date}</td>
    <td>${data.description}</td>
    <td>${formatCurrency(data.amount)}</td>
    <td>
      <button class="edit-btn">Edit</button>
      <button class="delete-btn">Delete</button>
    </td>
  `;

  if(type === "income") {
    totalIncome += data.amount;
  } else {
    totalExpense += data.amount;
  }

  row.querySelector(".delete-btn").onclick = () => {
    if (confirm("Bạn chắc chắn muốn xóa?")) {
      if (type === "income") totalIncome -= data.amount;
      else totalExpense -= data.amount;
      row.remove();
      updateTotals();
    }
  };

  row.querySelector(".edit-btn").onclick = () => {
    const newDate = prompt("Date:", data.date) || data.date;
    const newDesc = prompt("Description:", data.description) || data.description;
    const newAmount = parseInt(prompt("Amount:", data.amount) || data.amount);

    if (!isNaN(newAmount)) {
      if (type === "income") {
        totalIncome -= data.amount;
        totalIncome += newAmount;
      } else {
        totalExpense -= data.amount;
        totalExpense += newAmount;
      }

      data.date = newDate;
      data.description = newDesc;
      data.amount = newAmount;

      row.children[0].textContent = newDate;
      row.children[1].textContent = newDesc;
      row.children[2].textContent = formatCurrency(newAmount);
      updateTotals();
    }
  };

  tableBody.appendChild(row);
}

document.querySelectorAll(".tab").forEach((tab) => {
  tab.addEventListener("click", () => {
    document
      .querySelectorAll(".tab")
      .forEach((t) => t.classList.remove("active"));
    tab.classList.add("active");

    const target = tab.getAttribute("data-target");
    document
      .querySelectorAll(".tab-section")
      .forEach((sec) => sec.classList.remove("active"));
    document.getElementById(`${target}-section`).classList.add("active");
  });
});
