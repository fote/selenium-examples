require "selenium-webdriver"

#driver = Selenium::WebDriver.for :remote, desired_capabilities: :chrome

driver = Selenium::WebDriver.for :remote, url: "http://127.0.0.1:4444/wd/hub"

driver.navigate.to "http://google.com"

element = driver.find_element(name: 'q')
element.send_keys "Hello WebDriver!"
element.submit

puts driver.title

driver.quit
