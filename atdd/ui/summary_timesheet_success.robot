*** Settings ***
Library    SeleniumLibrary
Suite Setup    Open Browser    ${LOCAL_URL}    ${BROWSER} 
Test Setup    Go To    ${LOCAL_URL}
Suite Teardown    Close Browser

Resource    resource_success.robot

***Variables***
${URL}    http://localhost:8080/home

*** Test Case ***
ดูผลสรุปในหน้า PAYMENTS เดือน 12 ปีคริสต์ศักราช 2018 และเปลี่ยนสถานะการโอนเงินของ Prathan Dansakulcharoenkit เป็น ถูกต้อง สำเร็จ
    เข้าสู่ระบบ 
    ใส่เดือนและปีที่ต้องการดูสรุปผล    12\t2018    12-DECEMBER2018-TIMESHEET
    แสดงรายละเอียดของพนักงาน ID    0    Prathan Dansakulcharoenkit    ฿ 75,000.00    ฿ 30,000.00    ฿ 40,000.00    ฿ 145,000.00    ฿ 80,000.00    ฿ 5,000.00    ฿ 0.00    ฿ 75,000.00    ฿ 65,000.00    10%    ฿ 6,500.00    ฿ 58,500.00    ฿ 133,500.00    รอการตรวจสอบ    ${EMPTY}    ${EMPTY}
    แสดงรายละเอียดของพนักงาน ID    1    Prathan Dansakulcharoenkit    ฿ 0.00    ฿ 40,000.00    ฿ 0.00    ฿ 40,000.00    ฿ 0.00    ฿ 0.00    ฿ 0.00    ฿ 0.00    ฿ 40,000.00    10%    ฿ 4,000.00    ฿ 36,000.00    ฿ 36,000.00    รอการตรวจสอบ    ${EMPTY}    ${EMPTY} 
    แสดงรายละเอียดของพนักงาน ID    2    Nareenart Narunchon    ฿ 0.00    ฿ 0.00    ฿ 6,500.00    ฿ 6,500.00    ฿ 25,000.00    ฿ 0.00    ฿ 750.00    ฿ 24,250.00    ฿ 6,500.00    5%    ฿ 325.00    ฿ 6,175.00    ฿ 30,425.00    รอการตรวจสอบ    ${EMPTY}    ${EMPTY}  
    แสดงรายละเอียดของพนักงาน ID    3    Thawatchai Jongsuwanpaisan    ฿ 50,000.00    ฿ 70,000.00    ฿ 10,000.00    ฿ 130,000.00    ฿ 40,000.00    ฿ 5,000.00    ฿ 0.00    ฿ 35,000.00    ฿ 90,000.00    10%    ฿ 9,000.00    ฿ 81,000.00    ฿ 116,000.00    รอการตรวจสอบ    ${EMPTY}    ${EMPTY} 
    แสดงรายละเอียดของพนักงาน ID    4    Pruth Udompruge    ฿ 130,000.00    ฿ 0.00    ฿ 10,000.00    ฿ 140,000.00    ฿ 40,000.00    ฿ 5,000.00    ฿ 0.00    ฿ 35,000.00    ฿ 100,000.00    10%    ฿ 10,000.00    ฿ 90,000.00    ฿ 125,000.00    รอการตรวจสอบ    ${EMPTY}    ${EMPTY}
    เปลี่ยนสถานะการโอน วันที่โอน และหมายเหตุ    0    ถูกต้อง    28/12/2018    ค่าตั๋วที่ออกไปก่อน = 169,380.00 บาท
    แสดงรายละเอียดของพนักงาน ID    0    Prathan Dansakulcharoenkit    ฿ 75,000.00    ฿ 30,000.00    ฿ 40,000.00    ฿ 145,000.00    ฿ 80,000.00    ฿ 5,000.00    ฿ 0.00    ฿ 75,000.00    ฿ 65,000.00    10%    ฿ 6,500.00    ฿ 58,500.00    ฿ 133,500.00    ถูกต้อง    28/12/2018    ค่าตั๋วที่ออกไปก่อน = 169,380.00 บาท