export function daysLeft() {
    remaining_days = (Date(deadline).getTime() - Date.now()) / (1000 * 3600 * 24);
    return remaining_days.toFixed(0);
};

export function checkImage(url, callback) {
    const img = new Image();
    img.src = url;

    if (img.complete) callback(true);
    img.onload = () => callback(true);
    img.onerror = () => callback(false);
};

export function calculateBarPercentage(goal, raised) {
    return Math.round((raised * 100) / goal);
};
