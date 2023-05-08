#![no_main]
#![no_std]

#[panic_handler]
#[no_mangle]
pub fn panic(_info: &core::panic::PanicInfo) -> ! {
    loop {}
}

#[no_mangle]
extern "C" fn sub(a: i32, b: i32) -> i32 {
    a - b
}
