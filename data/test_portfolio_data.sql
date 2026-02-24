-- Portfolio 测试数据
-- 插入作品展示的测试数据

-- 插入作品数据
INSERT INTO portfolios (title, description, cover_url, project_type, status, demo_url, github_url, featured, sort_order, overview, role, duration, client, challenge, solution, created_at, updated_at) VALUES
-- 前端项目
('个人博客系统', '基于 Vue3 + Vite 构建的现代化个人博客系统，支持 Markdown 渲染、评论系统、标签分类等功能。', 'https://picsum.photos/seed/blog/800/600', 'frontend', 'completed', 'https://blog.example.com', 'https://github.com/example/blog', true, 1, '一个功能完整的个人博客系统，采用现代化技术栈构建。', '前端开发', '2个月', '个人', '需要实现一个高性能的博客系统，支持实时预览和 SEO 优化。', '使用 Vue3 + Vite + Pinia 构建，集成 md-editor-v3 实现 Markdown 渲染，后端 API 采用 Go + Gin。', NOW(), NOW()),

('在线代码编辑器', '支持多种编程语言的在线代码编辑器，具有语法高亮、代码补全、实时协作等功能。', 'https://picsum.photos/seed/editor/800/600', 'frontend', 'completed', 'https://code-editor.example.com', 'https://github.com/example/code-editor', true, 2, '基于 Monaco Editor 构建的在线代码编辑器。', '前端开发', '3个月', '开源项目', '实现一个性能优异的在线编辑器，支持大型文件编辑。', '使用 Web Worker 进行语法分析，虚拟滚动优化大文件性能。', NOW(), NOW()),

('任务管理看板', '类似 Trello 的可视化任务管理工具，支持拖拽排序、看板视图、团队协作等功能。', 'https://picsum.photos/seed/kanban/800/600', 'frontend', 'completed', 'https://kanban.example.com', 'https://github.com/example/kanban', false, 3, '轻量级任务管理工具，界面简洁美观。', '全栈开发', '4个月', '创业团队', '需要一个简单易用的项目管理工具。', 'Vue3 + Go + PostgreSQL，使用 dnd-kit 实现拖拽功能。', NOW(), NOW()),

('天气应用', '实时天气查询应用，支持全球城市搜索、7天天气预报、空气质量指数等功能。', 'https://picsum.photos/seed/weather/800/600', 'frontend', 'completed', 'https://weather.example.com', 'https://github.com/example/weather', false, 4, '精美的天气 UI 界面，支持深色模式。', '前端开发', '1个月', '个人项目', '学习天气 API 的使用。', '调用和风天气 API，使用 ECharts 绘制温度趋势图。', NOW(), NOW()),

('音乐播放器', '支持多种音频格式的网页音乐播放器，具有歌词同步、播放列表、可视化效果等功能。', 'https://picsum.photos/seed/music/800/600', 'frontend', 'developing', NULL, 'https://github.com/example/music-player', false, 5, '沉浸式音乐体验，支持在线音乐搜索。', '前端开发', '2个月', '个人项目', '想做一个有特色的音乐播放器。', '使用 Web Audio API，集成网易云音乐 API。', NOW(), NOW()),

-- VibeCoding 项目
('AI 代码助手插件', 'VS Code 插件，集成 GPT-4 进行代码补全、代码解释、重构建议等功能。', 'https://picsum.photos/seed/ai-assistant/800/600', 'vibecoding', 'completed', 'https://marketplace.visualstudio.com/items?itemName=example.ai-assistant', 'https://github.com/example/ai-assistant', true, 6, '提升编程效率的 AI 助手。', '全栈开发', '6个月', '个人', '探索 AI 辅助编程的可能性。', '使用 OpenAI API，基于 VS Code Extension API 开发。', NOW(), NOW()),

('智能图片标注工具', '基于计算机视觉的图片标注工具，支持物体检测、语义分割、图像分类等任务。', 'https://picsum.photos/seed/label-tool/800/600', 'vibecoding', 'completed', 'https://label-tool.example.com', 'https://github.com/example/label-tool', false, 7, '用于机器学习数据集标注的专业工具。', '全栈开发', '5个月', 'AI 研究团队', '需要高效的数据标注工具。', '使用 TensorFlow.js 进行模型推理，Vue3 构建界面。', NOW(), NOW()),

-- 全栈项目
('电商平台', '功能完整的 B2C 电商平台，包含商品管理、购物车、订单系统、支付集成等功能。', 'https://picsum.photos/seed/ecommerce/800/600', 'fullstack', 'completed', 'https://shop.example.com', 'https://github.com/example/shop', true, 8, '现代化电商解决方案，支持多商户入驻。', '全栈开发', '8个月', '客户项目', '客户需要一个可扩展的电商平台。', '后端 Go + gRPC，前端 Vue3 + Nuxt，使用 PostgreSQL + Redis。', NOW(), NOW()),

('实时聊天应用', '支持单人、群聊、文件传输的实时聊天应用，类似微信 Web 版。', 'https://picsum.photos/seed/chat/800/600', 'fullstack', 'completed', 'https://chat.example.com', 'https://github.com/example/chat', false, 9, '使用 WebSocket 实现低延迟通信。', '全栈开发', '4个月', '学习项目', '学习实时通信技术。', 'Go + WebSocket，前端 Vue3 + Vuex。', NOW(), NOW()),

('在线文档协作', '类似 Google Docs 的在线文档协作工具，支持多人实时编辑、评论、版本历史等功能。', 'https://picsum.photos/seed/docs/800/600', 'fullstack', 'developing', NULL, 'https://github.com/example/docs', false, 10, '基于 OT 算法的实时协作。', '全栈开发', '进行中', '团队内部工具', '团队需要一个共享文档编辑工具。', '使用 Yjs 实现 CRDT，后端 Go 处理权限控制。', NOW(), NOW()),

-- 小程序
('健身打卡小程序', '帮助用户记录健身打卡、制定训练计划、分享健身成果的微信小程序。', 'https://picsum.photos/seed/fitness/800/600', 'miniprogram', 'completed', NULL, 'https://github.com/example/fitness-miniprogram', false, 11, '累计用户 10万+，包含训练课程和社区功能。', '前端开发', '3个月', '客户项目', '客户需要一个健身相关的小程序。', '微信小程序原生开发，后端 Go + MySQL。', NOW(), NOW()),

('美食推荐小程序', '基于地理位置的美食推荐小程序，包含商家入驻、用户点评、优惠活动等功能。', 'https://picsum.photos/seed/food/800/600', 'miniprogram', 'completed', NULL, 'https://github.com/example/food-miniprogram', false, 12, '探索身边美食，发现特色餐厅。', '全栈开发', '4个月', '个人项目', '结合兴趣做的小程序。', 'uni-app 开发，后端 Go。', NOW(), NOW()),

('二手交易小程序', '校园二手交易平台，支持商品发布、即时聊天、线下交易等功能。', 'https://picsum.photos/seed/trade/800/600', 'miniprogram', 'archived', NULL, 'https://github.com/example/trade-miniprogram', false, 13, '服务校园用户超过 5000 人。', '全栈开发', '2个月', '大学项目', '毕业后项目已归档。', '微信小程序，后端 Node.js。', NOW(), NOW()),

-- APP
('冥想助手 APP', '帮助用户进行冥想练习的应用，包含引导音频、进度跟踪、社区分享等功能。', 'https://picsum.photos/seed/meditation/800/600', 'app', 'completed', 'https://apps.apple.com/app/meditation', 'https://github.com/example/meditation-app', false, 14, '帮助用户养成冥想习惯，放松身心。', '跨平台开发', '6个月', '个人项目', '想做一款健康类应用。', '使用 Flutter 开发，后端 Go。', NOW(), NOW()),

('记账 APP', '简洁美观的个人记账应用，支持分类统计、预算管理、数据导出等功能。', 'https://picsum.photos/seed/finance/800/600', 'app', 'completed', 'https://apps.apple.com/app/finance', 'https://github.com/example/finance-app', false, 15, '自动化记账，智能分类消费。', '全栈开发', '5个月', '个人项目', '自己的记账需求。', 'React Native + Go，使用 Plaid 集成银行 API。', NOW(), NOW()),

-- 其他
('技术博客主题', '为 Hexo 博客框架开发的自定义主题，具有独特的视觉设计和丰富的交互效果。', 'https://picsum.photos/seed/theme/800/600', 'other', 'completed', 'https://github.com/example/hexo-theme', 'https://github.com/example/hexo-theme', true, 16, 'Star 数 1000+，被数百名博主使用。', '前端开发', '持续维护', '开源项目', '为 Hexo 贡献主题。', '使用 Pug 模板引擎，SCSS 编写样式。', NOW(), NOW()),

('开源组件库', '常用的 Vue3 组件库，包含按钮、表单、图表等 30+ 个高质量组件。', 'https://picsum.photos/seed/components/800/600', 'other', 'completed', 'https://components.example.com', 'https://github.com/example/components', false, 17, 'TypeScript + Vite 构建，完整文档和示例。', '前端开发', '持续维护', '开源项目', '沉淀通用组件。', '使用 VitePress 构建文档网站。', NOW(), NOW()),

('数据可视化大屏', '为企业客户开发的数据可视化大屏系统，实时展示业务数据和 KPI。', 'https://picsum.photos/seed/dashboard/800/600', 'other', 'completed', NULL, NULL, false, 18, '支持拖拽配置，响应式布局。', '全栈开发', '3个月', '客户项目', '客户需要数据展示大屏。', 'ECharts + Vue3，后端 Go 处理实时数据推送。', NOW(), NOW()),

('命令行工具集', '一系列实用的命令行工具，包括文件批量处理、日志分析、性能监控等。', 'https://picsum.photos/seed/cli/800/600', 'other', 'completed', 'https://github.com/example/cli-tools', 'https://github.com/example/cli-tools', false, 19, '提高开发效率的实用工具集。', '后端开发', '持续维护', '开源项目', '解决日常开发中的痛点。', '使用 Go 和 Python 开发。', NOW(), NOW()),

('服务器监控面板', '轻量级服务器监控面板，支持 CPU、内存、磁盘、网络等资源监控。', 'https://picsum.photos/seed/monitor/800/600', 'other', 'completed', 'https://monitor.example.com', 'https://github.com/example/monitor', false, 20, '支持告警通知，历史数据查询。', '全栈开发', '2个月', '个人项目', '需要监控服务器状态。', 'Go 后端采集数据，WebSocket 推送实时状态。', NOW(), NOW());

-- 插入技术栈关联数据
INSERT INTO portfolio_technologies (portfolio_technologies, technology, created_at, updated_at) VALUES
-- 个人博客系统 (id=1)
(1, 'Vue3', NOW(), NOW()), (1, 'Vite', NOW(), NOW()), (1, 'Pinia', NOW(), NOW()), (1, 'TypeScript', NOW(), NOW()), (1, 'TailwindCSS', NOW(), NOW()), (1, 'Go', NOW(), NOW()), (1, 'Gin', NOW(), NOW()),
-- 在线代码编辑器 (id=2)
(2, 'Monaco Editor', NOW(), NOW()), (2, 'React', NOW(), NOW()), (2, 'Web Worker', NOW(), NOW()), (2, 'TypeScript', NOW(), NOW()), (2, 'Vite', NOW(), NOW()),
-- 任务管理看板 (id=3)
(3, 'Vue3', NOW(), NOW()), (3, 'Go', NOW(), NOW()), (3, 'Gin', NOW(), NOW()), (3, 'PostgreSQL', NOW(), NOW()), (3, 'DnD-Kit', NOW(), NOW()), (3, 'Redis', NOW(), NOW()),
-- 天气应用 (id=4)
(4, 'Vue3', NOW(), NOW()), (4, 'ECharts', NOW(), NOW()), (4, 'Axios', NOW(), NOW()), (4, 'SCSS', NOW(), NOW()), (4, 'API', NOW(), NOW()),
-- 音乐播放器 (id=5)
(5, 'Web Audio API', NOW(), NOW()), (5, 'Vue3', NOW(), NOW()), (5, 'NeteaseCloudAPI', NOW(), NOW()), (5, 'Canvas', NOW(), NOW()),
-- AI 代码助手插件 (id=6)
(6, 'VS Code API', NOW(), NOW()), ( 6, 'OpenAI API', NOW(), NOW()), ( 6, 'TypeScript', NOW(), NOW()), ( 6, 'Node.js', NOW(), NOW()), ( 6, 'GPT', NOW(), NOW()),
-- 智能图片标注工具 (id=7)
(7, 'TensorFlow.js', NOW(), NOW()), (7, 'Vue3', NOW(), NOW()), (7, 'Go', NOW(), NOW()), (7, 'OpenCV', NOW(), NOW()), (7, 'CV', NOW(), NOW()),
-- 电商平台 (id=8)
(8, 'Vue3', NOW(), NOW()), (8, 'Nuxt', NOW(), NOW()), (8, 'Go', NOW(), NOW()), (8, 'gRPC', NOW(), NOW()), (8, 'PostgreSQL', NOW(), NOW()), (8, 'Redis', NOW(), NOW()), (8, 'Stripe', NOW(), NOW()),
-- 实时聊天应用 (id=9)
(9, 'Go', NOW(), NOW()), (9, 'WebSocket', NOW(), NOW()), (9, 'Vue3', NOW(), NOW()), (9, 'Vuex', NOW(), NOW()), (9, 'Socket.IO', NOW(), NOW()),
-- 在线文档协作 (id=10)
(10, 'Yjs', NOW(), NOW()), (10, 'Go', NOW(), NOW()), (10, 'Vue3', NOW(), NOW()), (10, 'CRDT', NOW(), NOW()), (10, 'OT', NOW(), NOW()),
-- 健身打卡小程序 (id=11)
(11, '微信小程序', NOW(), NOW()), (11, 'Go', NOW(), NOW()), (11, 'MySQL', NOW(), NOW()), (11, '云开发', NOW(), NOW()),
-- 美食推荐小程序 (id=12)
(12, 'uni-app', NOW(), NOW()), (12, 'Go', NOW(), NOW()), (12, 'MongoDB', NOW(), NOW()), (12, 'LBS', NOW(), NOW()),
-- 二手交易小程序 (id=13)
(13, '微信小程序', NOW(), NOW()), (13, 'Node.js', NOW(), NOW()), (13, 'Express', NOW(), NOW()), (13, 'MongoDB', NOW(), NOW()),
-- 冥想助手 APP (id=14)
(14, 'Flutter', NOW(), NOW()), (14, 'Go', NOW(), NOW()), (14, 'PostgreSQL', NOW(), NOW()), (14, 'Audio', NOW(), NOW()), (14, 'Health', NOW(), NOW()),
-- 记账 APP (id=15)
(15, 'React Native', NOW(), NOW()), (15, 'Go', NOW(), NOW()), (15, 'Plaid', NOW(), NOW()), (15, 'Finance', NOW(), NOW()), (15, 'Mobile', NOW(), NOW()),
-- 技术博客主题 (id=16)
(16, 'Pug', NOW(), NOW()), (16, 'SCSS', NOW(), NOW()), (16, 'Hexo', NOW(), NOW()), (16, 'Node.js', NOW(), NOW()), (16, 'Theme', NOW(), NOW()),
-- 开源组件库 (id=17)
(17, 'Vue3', NOW(), NOW()), (17, 'TypeScript', NOW(), NOW()), (17, 'Vite', NOW(), NOW()), (17, 'VitePress', NOW(), NOW()), (17, 'Components', NOW(), NOW()),
-- 数据可视化大屏 (id=18)
(18, 'ECharts', NOW(), NOW()), (18, 'Vue3', NOW(), NOW()), (18, 'Go', NOW(), NOW()), (18, 'WebSocket', NOW(), NOW()), (18, 'Dashboard', NOW(), NOW()),
-- 命令行工具集 (id=19)
(19, 'Go', NOW(), NOW()), (19, 'Python', NOW(), NOW()), (19, 'Shell', NOW(), NOW()), (19, 'CLI', NOW(), NOW()), (19, 'DevTools', NOW(), NOW()),
-- 服务器监控面板 (id=20)
(20, 'Go', NOW(), NOW()), (20, 'WebSocket', NOW(), NOW()), (20, 'Vue3', NOW(), NOW()), (20, 'Monitoring', NOW(), NOW()), (20, 'Server', NOW(), NOW());
