-- ppopgipang-spine: seed data for all entity tables
-- Usage: mysql -u <user> -p <db_name> < seeds/all_entities_seed.sql

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

TRUNCATE TABLE `applications`;
TRUNCATE TABLE `job_postings`;
TRUNCATE TABLE `user_achievements`;
TRUNCATE TABLE `user_stamps`;
TRUNCATE TABLE `stamps`;
TRUNCATE TABLE `achievements`;
TRUNCATE TABLE `moderation_actions`;
TRUNCATE TABLE `content_reports`;
TRUNCATE TABLE `push_subscriptions`;
TRUNCATE TABLE `notifications`;
TRUNCATE TABLE `proposals`;
TRUNCATE TABLE `reviews`;
TRUNCATE TABLE `store_photos`;
TRUNCATE TABLE `store_opening_hours`;
TRUNCATE TABLE `store_facilities`;
TRUNCATE TABLE `store_analytics`;
TRUNCATE TABLE `stores`;
TRUNCATE TABLE `store_type`;
TRUNCATE TABLE `trade_chat_message`;
TRUNCATE TABLE `trade_chat_room`;
TRUNCATE TABLE `trades`;
TRUNCATE TABLE `user_store_stats`;
TRUNCATE TABLE `user_store_bookmarks`;
TRUNCATE TABLE `user_search_history`;
TRUNCATE TABLE `user_progress`;
TRUNCATE TABLE `user_loots`;
TRUNCATE TABLE `loot_likes`;
TRUNCATE TABLE `certification_reasons`;
TRUNCATE TABLE `certification_tags`;
TRUNCATE TABLE `certification_photos`;
TRUNCATE TABLE `certifications`;
TRUNCATE TABLE `checkin_reason_presets`;
TRUNCATE TABLE `loot_comment_presets`;
TRUNCATE TABLE `loot_tags`;
TRUNCATE TABLE `users`;

SET FOREIGN_KEY_CHECKS = 1;

INSERT INTO `store_type` (`id`, `name`, `description`) VALUES
  (1, '인형뽑기', '클래식 크레인 머신 중심 매장'),
  (2, '가챠샵', '캡슐토이와 피규어 상품 중심 매장'),
  (3, '오락실', '아케이드 게임과 뽑기 복합 매장');

INSERT INTO `users` (`id`, `email`, `kakaoId`, `nickname`, `profileImage`, `isAdmin`, `refreshToken`, `adminPassword`, `mannerTemp`, `createdAt`, `updatedAt`) VALUES
  (1, 'admin@ppopgipang.com', 'kakao_admin_1001', '운영자팡', 'profiles/admin.png', 1, 'refresh-token-admin', '$2a$10$admin.seed.hash.value', 36.8, '2026-01-01 09:00:00.000000', '2026-01-31 09:00:00.000000'),
  (2, 'user1@ppopgipang.com', 'kakao_user_1002', '집게왕', 'profiles/user1.png', 0, 'refresh-token-user1', NULL, 37.2, '2026-01-02 09:00:00.000000', '2026-01-31 12:30:00.000000'),
  (3, 'user2@ppopgipang.com', 'kakao_user_1003', '초코팬더', 'profiles/user2.png', 0, 'refresh-token-user2', NULL, 36.5, '2026-01-03 09:00:00.000000', '2026-01-30 20:10:00.000000'),
  (4, 'user3@ppopgipang.com', 'kakao_user_1004', '야간헌터', 'profiles/user3.png', 0, 'refresh-token-user3', NULL, 36.9, '2026-01-04 09:00:00.000000', '2026-02-02 21:50:00.000000'),
  (5, 'user4@ppopgipang.com', 'kakao_user_1005', '피규어덕후', 'profiles/user4.png', 0, 'refresh-token-user4', NULL, 37.4, '2026-01-05 09:00:00.000000', '2026-02-03 18:20:00.000000');

INSERT INTO `loot_tags` (`id`, `name`, `iconName`, `sortOrder`, `isActive`, `createdAt`) VALUES
  (1, '인형', 'tags/doll.png', 1, 1, '2026-01-01 10:00:00.000000'),
  (2, '피규어', 'tags/figure.png', 2, 1, '2026-01-01 10:00:00.000000'),
  (3, '키링', 'tags/keyring.png', 3, 1, '2026-01-01 10:00:00.000000'),
  (4, '캐릭터굿즈', 'tags/merch.png', 4, 1, '2026-01-01 10:00:00.000000'),
  (5, '한정판', 'tags/limited.png', 5, 1, '2026-01-01 10:00:00.000000');

INSERT INTO `checkin_reason_presets` (`id`, `content`, `sortOrder`, `isActive`, `createdAt`) VALUES
  (1, '오늘은 원하는 상품이 없었어요', 1, 1, '2026-01-01 11:00:00.000000'),
  (2, '기계 상태가 아쉬웠어요', 2, 1, '2026-01-01 11:00:00.000000'),
  (3, '다음에 다시 도전할게요', 3, 1, '2026-01-01 11:00:00.000000'),
  (4, '대기 인원이 너무 많았어요', 4, 1, '2026-01-01 11:00:00.000000'),
  (5, '예산을 다 써서 종료했어요', 5, 1, '2026-01-01 11:00:00.000000');

INSERT INTO `loot_comment_presets` (`id`, `content`, `sortOrder`, `isActive`, `createdAt`) VALUES
  (1, '한 번에 뽑아서 기분 최고!', 1, 1, '2026-01-01 11:10:00.000000'),
  (2, '집게 힘이 좋아서 금방 성공했어요', 2, 1, '2026-01-01 11:10:00.000000'),
  (3, '중반부터 감 잡고 역전 성공', 3, 1, '2026-01-01 11:10:00.000000'),
  (4, '친구랑 내기했는데 이겼어요', 4, 1, '2026-01-01 11:10:00.000000'),
  (5, '오늘 운이 좋아서 연속 득템!', 5, 1, '2026-01-01 11:10:00.000000');

INSERT INTO `job_postings` (`id`, `title`, `description`, `department`, `positionType`, `location`, `isActive`, `createdAt`, `updatedAt`) VALUES
  (1, '매장 운영 매니저', '오프라인 매장 운영 및 고객 응대', '운영팀', 'full-time', '서울 마포구', 1, '2026-01-05 09:00:00.000000', '2026-01-15 09:00:00.000000'),
  (2, '서비스 기획 인턴', '앱 기능 기획 및 운영 지표 분석', '프로덕트팀', 'intern', '서울 강남구', 1, '2026-01-10 09:00:00.000000', '2026-01-20 09:00:00.000000'),
  (3, '커뮤니티 운영 담당자', '유저 이벤트 운영 및 문의 대응', '커뮤니티팀', 'contract', '서울 송파구', 1, '2026-01-12 10:00:00.000000', '2026-01-28 09:30:00.000000'),
  (4, '백엔드 엔지니어', '거래/알림 서비스 고도화', '개발팀', 'full-time', '서울 강서구', 1, '2026-01-18 10:00:00.000000', '2026-02-01 11:00:00.000000');

INSERT INTO `stores` (`id`, `name`, `address`, `region1`, `region2`, `latitude`, `longitude`, `phone`, `averageRating`, `typeId`, `createdAt`, `updatedAt`) VALUES
  (1, '홍대 뽑기스팟', '서울 마포구 와우산로 12', '서울특별시', '마포구', 37.556712, 126.923441, '02-1111-2222', 4.6, 1, '2026-01-01 12:00:00.000000', '2026-01-31 22:00:00.000000'),
  (2, '강남 가챠랩', '서울 강남구 테헤란로 88', '서울특별시', '강남구', 37.501240, 127.039603, '02-3333-4444', 4.3, 2, '2026-01-02 12:00:00.000000', '2026-01-31 22:10:00.000000'),
  (3, '건대 크레인존', '서울 광진구 아차산로 55', '서울특별시', '광진구', 37.540652, 127.070712, '02-5555-6666', 4.1, 1, '2026-01-03 12:00:00.000000', '2026-02-02 21:00:00.000000'),
  (4, '잠실 플레이아레나', '서울 송파구 올림픽로 240', '서울특별시', '송파구', 37.511977, 127.098106, '02-7777-8888', 4.7, 3, '2026-01-04 12:00:00.000000', '2026-02-03 23:00:00.000000');

INSERT INTO `store_analytics` (`storeId`, `congestionScore`, `successProb`, `recentLootCount`, `hotTimeJson`, `lastAnalyzedAt`) VALUES
  (1, 61, 72, 9, '{"18":72,"19":75,"20":70}', '2026-02-10 20:00:00.000000'),
  (2, 48, 64, 5, '{"17":61,"18":64,"19":67}', '2026-02-10 20:00:00.000000'),
  (3, 57, 69, 7, '{"16":65,"17":68,"18":69}', '2026-02-10 20:00:00.000000'),
  (4, 73, 78, 13, '{"19":79,"20":81,"21":77}', '2026-02-10 20:00:00.000000');

INSERT INTO `store_facilities` (`storeId`, `machineCount`, `paymentMethods`, `notes`) VALUES
  (1, 34, '["cash","card","qr"]', '신형 집게 머신 비율이 높음'),
  (2, 22, '["card","qr"]', '가챠 전용 결제 키오스크 운영'),
  (3, 28, '["cash","card"]', '중형 기계 라인업이 다양함'),
  (4, 42, '["cash","card","qr"]', '야간 방문객이 많은 대형 매장');

INSERT INTO `store_opening_hours` (`id`, `storeId`, `dayOfWeek`, `openTime`, `closeTime`, `isClosed`) VALUES
  (1, 1, 1, '2026-01-05 11:00:00.000000', '2026-01-05 23:00:00.000000', 0),
  (2, 1, 2, '2026-01-06 11:00:00.000000', '2026-01-06 23:00:00.000000', 0),
  (3, 1, 3, '2026-01-07 11:00:00.000000', '2026-01-07 23:00:00.000000', 0),
  (4, 1, 4, '2026-01-08 11:00:00.000000', '2026-01-08 23:00:00.000000', 0),
  (5, 2, 1, '2026-01-05 10:30:00.000000', '2026-01-05 22:30:00.000000', 0),
  (6, 2, 2, '2026-01-06 10:30:00.000000', '2026-01-06 22:30:00.000000', 0),
  (7, 2, 3, '2026-01-07 10:30:00.000000', '2026-01-07 22:30:00.000000', 0),
  (8, 2, 4, '2026-01-08 10:30:00.000000', '2026-01-08 22:30:00.000000', 0),
  (9, 3, 5, '2026-01-09 12:00:00.000000', '2026-01-09 23:30:00.000000', 0),
  (10, 3, 6, '2026-01-10 12:00:00.000000', '2026-01-10 23:59:59.000000', 0),
  (11, 4, 5, '2026-01-09 10:00:00.000000', '2026-01-09 23:59:59.000000', 0),
  (12, 4, 6, '2026-01-10 10:00:00.000000', '2026-01-10 23:59:59.000000', 0);

INSERT INTO `store_photos` (`id`, `storeId`, `type`, `imageName`) VALUES
  (1, 1, 'cover', 'stores/1-cover.jpg'),
  (2, 1, 'inside', 'stores/1-inside.jpg'),
  (3, 1, 'sign', 'stores/1-sign.jpg'),
  (4, 2, 'cover', 'stores/2-cover.jpg'),
  (5, 2, 'sign', 'stores/2-sign.jpg'),
  (6, 2, 'inside', 'stores/2-inside.jpg'),
  (7, 3, 'cover', 'stores/3-cover.jpg'),
  (8, 3, 'roadview', 'stores/3-roadview.jpg'),
  (9, 3, 'inside', 'stores/3-inside.jpg'),
  (10, 4, 'cover', 'stores/4-cover.jpg'),
  (11, 4, 'inside', 'stores/4-inside.jpg'),
  (12, 4, 'sign', 'stores/4-sign.jpg');

INSERT INTO `user_progress` (`userId`, `level`, `exp`, `streakDays`, `lastActivityAt`) VALUES
  (1, 10, 2580, 14, '2026-02-12 22:00:00.000000'),
  (2, 6, 920, 5, '2026-02-12 20:10:00.000000'),
  (3, 4, 430, 2, '2026-02-11 21:05:00.000000'),
  (4, 8, 1460, 9, '2026-02-12 23:20:00.000000'),
  (5, 5, 770, 4, '2026-02-12 19:40:00.000000');

INSERT INTO `certifications` (`id`, `userId`, `storeId`, `type`, `occurredAt`, `latitude`, `longitude`, `exp`, `comment`, `rating`, `createdAt`) VALUES
  (1, 2, 1, 'loot', '2026-02-01 19:20:00.000000', 37.556700, 126.923400, 120, '원하던 인형 득템 성공!', NULL, '2026-02-01 19:25:00.000000'),
  (2, 3, 2, 'checkin', '2026-02-02 18:10:00.000000', 37.501200, 127.039600, 40, NULL, 'normal', '2026-02-02 18:12:00.000000'),
  (3, 4, 3, 'loot', '2026-02-03 21:05:00.000000', 37.540600, 127.070700, 140, '마지막 집게에서 극적으로 성공!', NULL, '2026-02-03 21:07:00.000000'),
  (4, 2, 2, 'checkin', '2026-02-04 17:55:00.000000', 37.501200, 127.039500, 35, NULL, 'bad', '2026-02-04 17:57:00.000000'),
  (5, 5, 1, 'loot', '2026-02-05 20:15:00.000000', 37.556710, 126.923420, 160, '신상 캐릭터 키링을 뽑았어요!', NULL, '2026-02-05 20:17:00.000000'),
  (6, 3, 4, 'checkin', '2026-02-06 22:05:00.000000', 37.511970, 127.098100, 45, NULL, 'good', '2026-02-06 22:07:00.000000'),
  (7, 4, 4, 'checkin', '2026-02-07 20:45:00.000000', 37.511980, 127.098110, 42, NULL, 'normal', '2026-02-07 20:47:00.000000'),
  (8, 5, 3, 'loot', '2026-02-08 16:25:00.000000', 37.540640, 127.070700, 150, '한정판 피규어 득템!', NULL, '2026-02-08 16:27:00.000000');

INSERT INTO `certification_photos` VALUES
  (1, 1, 'certifications/1-1.jpg', 1),
  (2, 1, 'certifications/1-2.jpg', 2),
  (3, 3, 'certifications/3-1.jpg', 1),
  (4, 5, 'certifications/5-1.jpg', 1),
  (5, 5, 'certifications/5-2.jpg', 2),
  (6, 8, 'certifications/8-1.jpg', 1),
  (7, 8, 'certifications/8-2.jpg', 2),
  (8, 8, 'certifications/8-3.jpg', 3);

INSERT INTO `certification_tags` VALUES
  (1, 1),
  (1, 2),
  (3, 2),
  (3, 4),
  (5, 1),
  (5, 5),
  (8, 3),
  (8, 5);

SET @cr_cert_col = (
  SELECT CASE
    WHEN EXISTS (
      SELECT 1
      FROM information_schema.COLUMNS
      WHERE TABLE_SCHEMA = DATABASE()
        AND TABLE_NAME = 'certification_reasons'
        AND COLUMN_NAME = 'certification_id'
    ) THEN 'certification_id'
    ELSE 'certificationId'
  END
);

SET @cr_reason_col = (
  SELECT CASE
    WHEN EXISTS (
      SELECT 1
      FROM information_schema.COLUMNS
      WHERE TABLE_SCHEMA = DATABASE()
        AND TABLE_NAME = 'certification_reasons'
        AND COLUMN_NAME = 'reason_id'
    ) THEN 'reason_id'
    WHEN EXISTS (
      SELECT 1
      FROM information_schema.COLUMNS
      WHERE TABLE_SCHEMA = DATABASE()
        AND TABLE_NAME = 'certification_reasons'
        AND COLUMN_NAME = 'checkin_reason_preset_id'
    ) THEN 'checkin_reason_preset_id'
    ELSE 'reasonId'
  END
);

SET @sql_certification_reasons = CONCAT(
  'INSERT INTO `certification_reasons` (`',
  @cr_cert_col,
  '`, `',
  @cr_reason_col,
  '`) VALUES (2, 1), (2, 3), (4, 2), (6, 4), (7, 5)'
);

PREPARE stmt_certification_reasons FROM @sql_certification_reasons;
EXECUTE stmt_certification_reasons;
DEALLOCATE PREPARE stmt_certification_reasons;

INSERT INTO `loot_likes` VALUES
  (1, 1, '2026-02-01 19:30:00.000000'),
  (3, 1, '2026-02-01 20:10:00.000000'),
  (4, 1, '2026-02-02 09:10:00.000000'),
  (2, 3, '2026-02-03 21:20:00.000000'),
  (5, 3, '2026-02-03 21:25:00.000000'),
  (2, 5, '2026-02-05 20:30:00.000000'),
  (3, 8, '2026-02-08 17:10:00.000000'),
  (1, 8, '2026-02-08 18:00:00.000000');

INSERT INTO `user_loots` (`id`, `userId`, `certificationId`, `title`, `category`, `estimatedPrice`, `rarity`, `aiConfidence`, `status`, `createdAt`) VALUES
  (1, 2, 1, '초코 곰돌이 대형 인형', '인형', 18000, 'rare', 0.94, 'selling', '2026-02-01 19:26:00.000000'),
  (2, 3, NULL, '한정판 미니 키링', '키링', 7000, 'common', 0.81, 'kept', '2026-02-03 14:10:00.000000'),
  (3, 4, 3, '하늘색 여우 피규어', '피규어', 26000, 'rare', 0.96, 'kept', '2026-02-03 21:08:00.000000'),
  (4, 5, 5, '레트로 곰 키링', '키링', 9000, 'common', 0.87, 'selling', '2026-02-05 20:18:00.000000'),
  (5, 5, 8, '한정판 아스트로 피규어', '피규어', 52000, 'legend', 0.98, 'sold', '2026-02-08 16:30:00.000000'),
  (6, 2, NULL, '교환용 소형 인형 세트', '인형', 12000, 'common', 0.73, 'exchanged', '2026-02-09 13:40:00.000000');

INSERT INTO `user_search_history` (`id`, `userId`, `keyword`, `searchedAt`) VALUES
  (1, 2, '홍대 뽑기', '2026-02-10 13:20:00.000000'),
  (2, 2, '강남 가챠', '2026-02-11 14:00:00.000000'),
  (3, 3, '피규어 매장', '2026-02-11 20:40:00.000000'),
  (4, 4, '건대 인형뽑기', '2026-02-12 12:30:00.000000'),
  (5, 5, '잠실 오락실', '2026-02-12 15:10:00.000000'),
  (6, 3, '한정판 키링', '2026-02-12 21:20:00.000000'),
  (7, 4, '늦게까지 하는 뽑기방', '2026-02-12 22:05:00.000000'),
  (8, 5, '거래 게시판 시세', '2026-02-12 23:10:00.000000');

INSERT INTO `user_store_bookmarks` (`id`, `userId`, `storeId`, `createdAt`) VALUES
  (1, 2, 1, '2026-02-01 18:00:00.000000'),
  (2, 3, 2, '2026-02-02 17:40:00.000000'),
  (3, 4, 3, '2026-02-03 20:00:00.000000'),
  (4, 5, 4, '2026-02-04 19:10:00.000000'),
  (5, 2, 4, '2026-02-06 22:30:00.000000'),
  (6, 3, 1, '2026-02-07 11:15:00.000000');

INSERT INTO `user_store_stats` (`userId`, `storeId`, `visitCount`, `lootCount`, `lastVisitedAt`, `tier`) VALUES
  (2, 1, 6, 2, '2026-02-12 19:50:00.000000', 'master'),
  (3, 2, 3, 0, '2026-02-12 18:10:00.000000', 'visited'),
  (1, 1, 1, 0, '2026-02-08 15:00:00.000000', 'visited'),
  (4, 3, 8, 3, '2026-02-12 22:40:00.000000', 'master'),
  (5, 4, 5, 2, '2026-02-12 21:55:00.000000', 'visited'),
  (2, 4, 2, 0, '2026-02-11 22:00:00.000000', 'visited'),
  (3, 1, 1, 1, '2026-02-10 20:20:00.000000', 'visited');

INSERT INTO `reviews` (`id`, `rating`, `content`, `images`, `createdAt`, `updatedAt`, `userId`, `storeId`) VALUES
  (1, 5, '기계 상태가 좋아서 재방문 의사 있음', '["reviews/1-1.jpg"]', '2026-02-03 20:00:00.000000', '2026-02-03 20:00:00.000000', 2, 1),
  (2, 4, '직원 응대가 친절하고 매장도 깔끔해요', '[]', '2026-02-04 18:30:00.000000', '2026-02-04 18:30:00.000000', 3, 2),
  (3, 4, '야간에도 사람이 많지만 회전이 빨라요', '["reviews/3-1.jpg","reviews/3-2.jpg"]', '2026-02-05 23:10:00.000000', '2026-02-05 23:10:00.000000', 4, 4),
  (4, 5, '한정판 상품 입고가 빨라요', '["reviews/4-1.jpg"]', '2026-02-06 16:20:00.000000', '2026-02-06 16:20:00.000000', 5, 3),
  (5, 3, '주말엔 대기가 길어요', '[]', '2026-02-08 14:00:00.000000', '2026-02-08 14:00:00.000000', 2, 4),
  (6, 5, '관리 상태 최고였어요', '["reviews/6-1.jpg"]', '2026-02-09 19:30:00.000000', '2026-02-09 19:30:00.000000', 5, 1);

INSERT INTO `proposals` (`id`, `name`, `address`, `latitude`, `longitude`, `status`, `createdAt`, `userId`, `storeId`) VALUES
  (1, '건대 뽑기존', '서울 광진구 동일로 20', 37.540200, 127.069000, 'pending', '2026-02-05 13:00:00.000000', 2, NULL),
  (2, '잠실 가챠센터', '서울 송파구 올림픽로 100', 37.514100, 127.102200, 'approved', '2026-02-06 15:20:00.000000', 3, 2),
  (3, '신촌 크레인하우스', '서울 서대문구 신촌로 40', 37.556000, 126.938000, 'rejected', '2026-02-07 12:00:00.000000', 4, NULL),
  (4, '합정 플레이존', '서울 마포구 양화로 45', 37.549600, 126.913600, 'pending', '2026-02-09 11:30:00.000000', 5, NULL);

INSERT INTO `trades` (`id`, `title`, `description`, `images`, `price`, `type`, `status`, `createdAt`, `updatedAt`, `userId`, `lootId`) VALUES
  (1, '초코 곰돌이 판매', '상태 A급, 직거래 선호', '["trades/1-1.jpg","trades/1-2.jpg"]', 22000, 'sale', 'active', '2026-02-07 11:00:00.000000', '2026-02-10 09:00:00.000000', 2, 1),
  (2, '키링 교환 구해요', '고양이 키링과 교환 원합니다', '["trades/2-1.jpg"]', NULL, 'exchange', 'reserved', '2026-02-08 16:30:00.000000', '2026-02-11 14:20:00.000000', 3, 2),
  (3, '여우 피규어 팝니다', '박스 포함, 하자 없음', '["trades/3-1.jpg","trades/3-2.jpg"]', 31000, 'sale', 'completed', '2026-02-09 15:20:00.000000', '2026-02-11 18:00:00.000000', 4, 3),
  (4, '레트로 키링 판매', '두 개 이상 구매 시 네고 가능', '["trades/4-1.jpg"]', 8500, 'sale', 'active', '2026-02-10 19:10:00.000000', '2026-02-12 09:00:00.000000', 5, 4),
  (5, '인형 세트 교환', '피규어로 교환 원해요', '["trades/5-1.jpg"]', NULL, 'exchange', 'cancelled', '2026-02-11 13:10:00.000000', '2026-02-12 10:10:00.000000', 2, 6);

INSERT INTO `trade_chat_room` (`id`, `tradeId`, `sellerId`, `buyerId`, `createdAt`, `updatedAt`) VALUES
  (1, 1, 2, 3, '2026-02-07 11:30:00.000000', '2026-02-07 11:40:00.000000'),
  (2, 2, 3, 1, '2026-02-08 17:00:00.000000', '2026-02-08 17:10:00.000000'),
  (3, 3, 4, 5, '2026-02-09 16:00:00.000000', '2026-02-10 09:30:00.000000'),
  (4, 4, 5, 2, '2026-02-10 20:00:00.000000', '2026-02-12 08:40:00.000000');

INSERT INTO `trade_chat_message` (`id`, `roomId`, `senderId`, `message`, `imageName`, `isRead`, `sentAt`) VALUES
  (1, 1, 3, '혹시 오늘 저녁 거래 가능할까요?', NULL, 1, '2026-02-07 11:31:00.000000'),
  (2, 1, 2, '네, 7시 홍대입구역 괜찮습니다.', NULL, 1, '2026-02-07 11:33:00.000000'),
  (3, 2, 1, '교환 가능한 키링 사진 보내드립니다.', 'chat/room2-1.jpg', 0, '2026-02-08 17:05:00.000000'),
  (4, 2, 3, '확인했습니다. 조건 괜찮아요.', NULL, 0, '2026-02-08 17:08:00.000000'),
  (5, 3, 5, '피규어 실물 사진 더 볼 수 있을까요?', NULL, 1, '2026-02-09 16:10:00.000000'),
  (6, 3, 4, '네, 방금 추가 업로드했습니다.', 'chat/room3-1.jpg', 1, '2026-02-09 16:12:00.000000'),
  (7, 3, 5, '좋습니다. 구매할게요.', NULL, 1, '2026-02-09 16:15:00.000000'),
  (8, 4, 2, '오늘 밤 직거래 가능해요?', NULL, 1, '2026-02-10 20:03:00.000000'),
  (9, 4, 5, '가능해요. 잠실역 9시 어떠세요?', NULL, 1, '2026-02-10 20:05:00.000000'),
  (10, 4, 2, '좋습니다. 그때 봬요.', NULL, 0, '2026-02-10 20:06:00.000000');

INSERT INTO `notifications` (`id`, `userId`, `type`, `title`, `message`, `payload`, `isRead`, `createdAt`) VALUES
  (1, 2, 'trade_msg', '새 채팅이 도착했어요', '거래 채팅방에 새 메시지가 있습니다.', '{"route":"/trades/chat/1","roomId":1}', 0, '2026-02-07 11:32:00.000000'),
  (2, 3, 'level_up', '레벨 업!', '레벨 4를 달성했습니다.', '{"route":"/mypage","level":4}', 1, '2026-02-09 08:00:00.000000'),
  (3, 1, 'store_hot', '혼잡도 알림', '홍대 뽑기스팟이 현재 붐비고 있어요.', '{"route":"/stores/1","storeId":1}', 0, '2026-02-10 18:20:00.000000'),
  (4, 4, 'badge', '새 배지 획득', '첫 득템 배지를 획득했습니다.', '{"route":"/gamification/badges","badge":"FIRST_LOOT"}', 0, '2026-02-03 21:30:00.000000'),
  (5, 5, 'trade_msg', '거래 문의가 왔어요', '레트로 키링 게시글에 새 문의가 있습니다.', '{"route":"/trades/chat/4","roomId":4}', 0, '2026-02-10 20:05:00.000000'),
  (6, 2, 'event', '주말 이벤트 안내', '주말 더블 포인트 이벤트가 시작됩니다.', '{"route":"/events/weekend-double"}', 1, '2026-02-11 09:00:00.000000'),
  (7, 3, 'proposal', '제안 결과 안내', '매장 제안이 승인되었습니다.', '{"route":"/proposals/2","status":"approved"}', 0, '2026-02-06 16:00:00.000000'),
  (8, 5, 'review', '리뷰 좋아요', '작성한 리뷰에 좋아요 10개가 달렸어요.', '{"route":"/reviews/6"}', 0, '2026-02-12 19:45:00.000000');

INSERT INTO `push_subscriptions` (`id`, `userId`, `platform`, `endpoint`, `authKey`, `p256dhKey`, `createdAt`) VALUES
  (1, 2, 'web', 'https://push.example.com/subscriptions/u2-web', 'auth-key-u2', 'p256dh-u2', '2026-02-01 09:00:00.000000'),
  (2, 3, 'android', 'https://fcm.googleapis.com/fcm/send/u3token', 'auth-key-u3', 'p256dh-u3', '2026-02-01 09:10:00.000000'),
  (3, 4, 'ios', 'https://api.push.apple.com/3/device/u4token', 'auth-key-u4', 'p256dh-u4', '2026-02-02 08:20:00.000000'),
  (4, 5, 'web', 'https://push.example.com/subscriptions/u5-web', 'auth-key-u5', 'p256dh-u5', '2026-02-02 09:30:00.000000'),
  (5, 1, 'android', 'https://fcm.googleapis.com/fcm/send/admin-token', 'auth-key-admin', 'p256dh-admin', '2026-02-03 07:00:00.000000');

INSERT INTO `content_reports` (`id`, `reporterId`, `targetType`, `targetId`, `reason`, `description`, `status`, `createdAt`) VALUES
  (1, 3, 'trade', 1, '허위 매물 의심', '상품 설명과 실제 상태가 다르다는 제보', 'open', '2026-02-11 10:00:00.000000'),
  (2, 2, 'chat', 3, '부적절한 메시지', '과도한 비매너 표현 사용', 'resolved', '2026-02-11 10:30:00.000000'),
  (3, 5, 'certification', 3, '사진 도용 의심', '외부 커뮤니티에서 본 사진과 유사합니다.', 'rejected', '2026-02-11 12:10:00.000000'),
  (4, 4, 'store', 2, '매장 정보 오류', '영업시간 정보가 실제와 다릅니다.', 'open', '2026-02-12 09:20:00.000000');

INSERT INTO `moderation_actions` (`id`, `reportId`, `adminId`, `action`, `note`, `createdAt`) VALUES
  (1, 2, 1, 'warn_user', '채팅 가이드 위반으로 경고 처리', '2026-02-11 11:00:00.000000'),
  (2, 3, 1, 'dismiss_report', '증빙 불충분으로 반려', '2026-02-11 13:00:00.000000'),
  (3, 4, 1, 'request_store_update', '운영팀에 매장 정보 재검증 요청', '2026-02-12 10:10:00.000000');

INSERT INTO `achievements` (`id`, `code`, `name`, `description`, `conditionJson`, `badgeImageName`, `isHidden`) VALUES
  (1, 'FIRST_LOOT', '첫 득템', '처음으로 득템 인증을 완료했어요.', '{"type":"loot_count","target":1}', 'badges/first_loot.png', 0),
  (2, 'CHECKIN_10', '성실 방문러', '체크인 인증 10회를 달성했어요.', '{"type":"checkin_count","target":10}', 'badges/checkin_10.png', 0),
  (3, 'TRADE_MASTER', '거래 달인', '거래 완료 5회를 달성했어요.', '{"type":"trade_completed","target":5}', 'badges/trade_master.png', 1),
  (4, 'REVIEWER_5', '리뷰어', '리뷰 5개를 작성했어요.', '{"type":"review_count","target":5}', 'badges/reviewer_5.png', 0),
  (5, 'NIGHT_HUNTER', '야간 헌터', '22시 이후 인증 3회를 달성했어요.', '{"type":"night_cert_count","target":3}', 'badges/night_hunter.png', 0);

INSERT INTO `stamps` (`id`, `storeId`, `imageName`) VALUES
  (1, 1, 'stamps/hongdae.png'),
  (2, 2, 'stamps/gangnam.png'),
  (3, 3, 'stamps/konkuk.png'),
  (4, 4, 'stamps/jamsil.png');

INSERT INTO `user_stamps` (`userId`, `stampId`, `acquiredAt`) VALUES
  (2, 1, '2026-02-01 19:25:00.000000'),
  (3, 2, '2026-02-02 18:12:00.000000'),
  (4, 3, '2026-02-03 21:07:00.000000'),
  (5, 1, '2026-02-05 20:17:00.000000'),
  (5, 3, '2026-02-08 16:27:00.000000'),
  (3, 4, '2026-02-06 22:07:00.000000');

INSERT INTO `user_achievements` (`userId`, `achievementId`, `earnedAt`) VALUES
  (2, 1, '2026-02-01 19:26:00.000000'),
  (3, 2, '2026-02-12 20:00:00.000000'),
  (4, 1, '2026-02-03 21:08:00.000000'),
  (5, 1, '2026-02-05 20:18:00.000000'),
  (4, 5, '2026-02-12 23:00:00.000000');

INSERT INTO `applications` (`id`, `jobPosting_id`, `name`, `email`, `phone`, `resumeName`, `memo`, `status`, `createdAt`) VALUES
  (1, 1, '김집게', 'applicant1@example.com', '010-1234-5678', 'resume-kim.pdf', '주말 근무 가능합니다.', 'new', '2026-02-06 09:30:00.000000'),
  (2, 2, '박가챠', 'applicant2@example.com', '010-9999-1111', 'resume-park.pdf', '데이터 분석 경험이 있습니다.', 'reviewing', '2026-02-07 10:10:00.000000'),
  (3, 3, '최운영', 'applicant3@example.com', '010-2222-3333', 'resume-choi.pdf', '커뮤니티 운영 경험 2년', 'new', '2026-02-08 11:40:00.000000'),
  (4, 4, '이백엔드', 'applicant4@example.com', '010-4444-5555', 'resume-lee.pdf', 'Go 기반 API 개발 경험 보유', 'interview', '2026-02-09 14:00:00.000000'),
  (5, 1, '정매니저', 'applicant5@example.com', '010-6666-7777', 'resume-jung.pdf', '오프라인 매장 오퍼레이션 경험 다수', 'accepted', '2026-02-10 09:50:00.000000');
