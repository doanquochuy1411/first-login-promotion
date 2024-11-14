**Quyết định Kỹ thuật**
Cơ sở dữ liệu: PostgreSQL
PostgreSQL được chọn vì tính ổn định, khả năng hỗ trợ các truy vấn phức tạp và khả năng mở rộng, giúp xử lý các tập dữ liệu lớn và các giao dịch phức tạp.

Framework Backend: Go & Gin
Go (Golang) được sử dụng vì hiệu suất cao, sự hiệu quả và đơn giản trong việc xây dựng các dịch vụ web có thể mở rộng. Gin là một framework web viết bằng Go, cung cấp một cách tiếp cận nhanh chóng và tối giản để xây dựng các API RESTful.

**Các Giả định**
 - Cài đặt và cấu hình PostgreSQL trên hệ thống nơi ứng dụng đang chạy.
 - Tệp dev.env chứa các cấu hình môi trường cần thiết, bao gồm chuỗi kết nối cơ sở dữ liệu, các khóa API và các biến khác cho môi trường phát triển cục bộ.
 - Ứng dụng dựa vào cấu trúc được cung cấp trong sơ đồ cơ sở dữ liệu PostgreSQL và tất cả các bảng và mối quan hệ phải được thiết lập đúng cách dựa trên cấu hình có trong tệp dev.env.

**Cải Tiến Trong Tương Lai**
Cải thiện Xử lý Lỗi: Cải thiện cơ chế xử lý lỗi để đảm bảo ứng dụng có thể phục hồi một cách mượt mà khi gặp sự cố kết nối cơ sở dữ liệu hoặc các lỗi runtime khác.
Cơ chế Caching: Triển khai caching để tối ưu hóa dữ liệu thường xuyên được truy cập, chẳng hạn như hồ sơ người dùng và trạng thái chiến dịch khuyến mãi, giảm tải cho cơ sở dữ liệu và cải thiện thời gian phản hồi.
Giao diện Quản lý Chiến dịch: Tạo một giao diện admin để quản lý các chiến dịch khuyến mãi, theo dõi thống kê sử dụng và giám sát hiệu suất chiến dịch.
Cảnh báo Hết Hạn Voucher: Triển khai hệ thống tự động cảnh báo người dùng khi voucher của họ gần hết hạn, giúp cải thiện trải nghiệm người dùng và tăng cường sự tham gia.
Phân tích Nâng Cao: Tích hợp các công cụ phân tích nâng cao và cơ chế theo dõi để giám sát sự thành công của các chiến dịch khuyến mãi và thu thập thông tin về hành vi người dùng.

**Hướng Dẫn Cài Đặt Cục Bộ**
Cài Đặt PostgreSQL:

Đảm bảo PostgreSQL đã được cài đặt và đang chạy trên máy cục bộ của bạn.
Tạo một cơ sở dữ liệu mới cho ứng dụng.
Cài Đặt Cơ Sở Dữ Liệu:

Sử dụng sơ đồ cơ sở dữ liệu có trong tệp đính kèm (schema.sql) để thiết lập các bảng cần thiết trong cơ sở dữ liệu PostgreSQL.
Tạo Tệp dev.env:

Tạo tệp có tên dev.env trong thư mục gốc của dự án.
Thêm nội dung sau vào tệp dev.env, điều chỉnh các giá trị phù hợp với thiết lập của bạn:
bash


### Hướng dẫn cài đặt

1. **Cài đặt PostgreSQL:**  
   Đảm bảo PostgreSQL đã được cài đặt và chạy trên máy của bạn. Tạo một cơ sở dữ liệu mới và lưu các thông tin kết nối.

2. **Tạo tệp `dev.env`:**  
   Tạo tệp có tên `dev.env` trong thư mục gốc của dự án và thêm các giá trị sau:
   
SERVER_ADDRESS = your_host
DB_SOURCE = postgres://username:password@localhost:5432/your_database_name
API_KEY = your_api_key


Thay thế username, password, và your_database_name bằng thông tin kết nối của bạn.

Chạy Ứng Dụng:

Đảm bảo tất cả các phụ thuộc đã được cài đặt và cấu hình (ví dụ: Gin, PostgreSQL driver cho Go).
Chạy ứng dụng Go:
bash
Sao chép mã
go run main.go
Ứng dụng sẽ chạy cục bộ trên http://localhost:8080.
Cải Tiến Cho Hệ Thống Chiến Dịch Khuyến Mãi
Để thiết kế và triển khai hệ thống chiến dịch khuyến mãi cho ứng dụng Trinity, có thể thực hiện những cải tiến sau:

**Logic Tạo Voucher:**

Khi người dùng đăng ký thông qua liên kết chiến dịch, tạo voucher duy nhất để áp dụng giảm giá 30% cho gói đăng ký Silver. Voucher này phải được liên kết với người dùng và có ngày hết hạn.
Lưu các voucher này trong một bảng vouchers riêng trong cơ sở dữ liệu, với các cột như user_id, voucher_code, discount_percentage, expiration_date và status.
Giới Hạn Chiến Dịch Cho 100 Người Dùng:

Tạo cơ chế theo dõi số lượng người dùng đã đăng ký thông qua liên kết chiến dịch.
Sau khi 100 người dùng đầu tiên đăng ký thành công, tắt chiến dịch hoặc hiển thị thông báo rằng chương trình khuyến mãi đã kết thúc.
Logic Hết Hạn Voucher:

Bao gồm logic tự động vô hiệu hóa voucher sau một khoảng thời gian xác định (ví dụ: 30 ngày kể từ ngày tạo voucher).
Đảm bảo rằng voucher hết hạn không thể sử dụng và cung cấp thông báo hết hạn cho người dùng.
Theo Dõi Việc Sử Dụng Voucher:

Thêm cơ chế theo dõi việc sử dụng voucher để giám sát xem voucher nào đã được sử dụng, và khi nào voucher đó được áp dụng.
Tích hợp với hệ thống đăng ký để chỉ áp dụng giảm giá nếu voucher hợp lệ và chưa hết hạn.
Thông Báo Người Dùng:

Thông báo cho người dùng khi họ nhận được voucher thành công.
Gửi nhắc nhở cho người dùng trước khi voucher hết hạn, khuyến khích họ sử dụng voucher trước khi nó hết hạn.
Bằng cách tập trung vào các cải tiến này, bạn có thể tạo ra một hệ thống chiến dịch khuyến mãi hiệu quả và thân thiện với người dùng cho ứng dụng Trinity, giúp tăng cường sự tham gia của người dùng và thu hút thêm nhiều đăng ký cho gói đăng ký Silver.
