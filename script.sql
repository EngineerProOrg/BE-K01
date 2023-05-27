USE [FSE]
GO
/****** Object:  Table [dbo].[class]    Script Date: 5/28/2023 12:01:12 AM ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[class](
	[class_id] [int] NOT NULL,
	[class_name] [varchar](255) NULL,
	[prof_id] [int] NULL,
	[course_id] [int] NULL,
	[room_id] [int] NULL,
PRIMARY KEY CLUSTERED 
(
	[class_id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY]
GO
/****** Object:  Table [dbo].[course]    Script Date: 5/28/2023 12:01:12 AM ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[course](
	[course_id] [int] NOT NULL,
	[course_name] [varchar](255) NULL,
PRIMARY KEY CLUSTERED 
(
	[course_id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY]
GO
/****** Object:  Table [dbo].[enroll]    Script Date: 5/28/2023 12:01:12 AM ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[enroll](
	[stud_id] [int] NOT NULL,
	[class_id] [int] NOT NULL,
	[grade] [varchar](3) NULL,
 CONSTRAINT [enroll_composite_key] PRIMARY KEY CLUSTERED 
(
	[stud_id] ASC,
	[class_id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY]
GO
/****** Object:  Table [dbo].[progessor]    Script Date: 5/28/2023 12:01:12 AM ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[progessor](
	[prof_id] [int] NOT NULL,
	[prof_lname] [varchar](50) NULL,
	[prof_fname] [varchar](50) NULL,
PRIMARY KEY CLUSTERED 
(
	[prof_id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY]
GO
/****** Object:  Table [dbo].[room]    Script Date: 5/28/2023 12:01:12 AM ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[room](
	[room_id] [int] NOT NULL,
	[room_loc] [varchar](50) NULL,
	[room_cap] [varchar](50) NULL,
	[class_id] [int] NULL,
PRIMARY KEY CLUSTERED 
(
	[room_id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY]
GO
/****** Object:  Table [dbo].[student]    Script Date: 5/28/2023 12:01:12 AM ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[student](
	[stud_id] [int] NOT NULL,
	[stud_fname] [varchar](50) NULL,
	[stud_lname] [varchar](50) NULL,
	[stud_street] [varchar](255) NULL,
	[stud_city] [varchar](50) NULL,
	[stud_zip] [varchar](10) NULL,
PRIMARY KEY CLUSTERED 
(
	[stud_id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY]
GO
INSERT [dbo].[class] ([class_id], [class_name], [prof_id], [course_id], [room_id]) VALUES (1, N'se1320', 1, 1, 1)
INSERT [dbo].[class] ([class_id], [class_name], [prof_id], [course_id], [room_id]) VALUES (2, N'se1322', 1, 2, 2)
INSERT [dbo].[class] ([class_id], [class_name], [prof_id], [course_id], [room_id]) VALUES (3, N'se1323', 1, 2, 3)
INSERT [dbo].[class] ([class_id], [class_name], [prof_id], [course_id], [room_id]) VALUES (4, N'se1324', 2, 2, 3)
GO
INSERT [dbo].[course] ([course_id], [course_name]) VALUES (1, N'java')
INSERT [dbo].[course] ([course_id], [course_name]) VALUES (2, N'c/c++')
INSERT [dbo].[course] ([course_id], [course_name]) VALUES (3, N'golang')
INSERT [dbo].[course] ([course_id], [course_name]) VALUES (4, N'javascript')
GO
INSERT [dbo].[enroll] ([stud_id], [class_id], [grade]) VALUES (1, 1, N'A')
INSERT [dbo].[enroll] ([stud_id], [class_id], [grade]) VALUES (1, 2, N'A')
INSERT [dbo].[enroll] ([stud_id], [class_id], [grade]) VALUES (2, 1, N'A')
INSERT [dbo].[enroll] ([stud_id], [class_id], [grade]) VALUES (3, 1, N'B')
INSERT [dbo].[enroll] ([stud_id], [class_id], [grade]) VALUES (3, 2, N'A')
GO
INSERT [dbo].[progessor] ([prof_id], [prof_lname], [prof_fname]) VALUES (1, N'prof', N'A')
INSERT [dbo].[progessor] ([prof_id], [prof_lname], [prof_fname]) VALUES (2, N'prof', N'B')
GO
INSERT [dbo].[room] ([room_id], [room_loc], [room_cap], [class_id]) VALUES (1, N'Building A, Room 101', N'30', 1)
INSERT [dbo].[room] ([room_id], [room_loc], [room_cap], [class_id]) VALUES (2, N'Building B, Room 102', N'40', 2)
INSERT [dbo].[room] ([room_id], [room_loc], [room_cap], [class_id]) VALUES (3, N'Building C, Room 103', N'30', 2)
INSERT [dbo].[room] ([room_id], [room_loc], [room_cap], [class_id]) VALUES (4, N'Building D, Room 104', N'50', 3)
GO
INSERT [dbo].[student] ([stud_id], [stud_fname], [stud_lname], [stud_street], [stud_city], [stud_zip]) VALUES (1, N'Tran', N'Quang Anh', N'Ha noi', N'Ha noi', N'100000')
INSERT [dbo].[student] ([stud_id], [stud_fname], [stud_lname], [stud_street], [stud_city], [stud_zip]) VALUES (2, N'Nguyen', N'Van A', N'Cau Giay', N'Ha noi', N'100000')
INSERT [dbo].[student] ([stud_id], [stud_fname], [stud_lname], [stud_street], [stud_city], [stud_zip]) VALUES (3, N'Nguyen', N'Van B', N'Dong Da', N'HA noi', N'100000')
INSERT [dbo].[student] ([stud_id], [stud_fname], [stud_lname], [stud_street], [stud_city], [stud_zip]) VALUES (4, N'Nguyen', N'Van C', N'Dong Da', N'HA noi', N'100000')
INSERT [dbo].[student] ([stud_id], [stud_fname], [stud_lname], [stud_street], [stud_city], [stud_zip]) VALUES (5, N'Nguyen', N'Van D', N'Dong Da', N'HA noi', N'100000')
GO
ALTER TABLE [dbo].[class]  WITH CHECK ADD FOREIGN KEY([course_id])
REFERENCES [dbo].[course] ([course_id])
GO
ALTER TABLE [dbo].[class]  WITH CHECK ADD FOREIGN KEY([prof_id])
REFERENCES [dbo].[progessor] ([prof_id])
GO
ALTER TABLE [dbo].[class]  WITH CHECK ADD FOREIGN KEY([room_id])
REFERENCES [dbo].[room] ([room_id])
GO
ALTER TABLE [dbo].[enroll]  WITH CHECK ADD FOREIGN KEY([class_id])
REFERENCES [dbo].[class] ([class_id])
GO
ALTER TABLE [dbo].[enroll]  WITH CHECK ADD FOREIGN KEY([stud_id])
REFERENCES [dbo].[student] ([stud_id])
GO
ALTER TABLE [dbo].[room]  WITH CHECK ADD FOREIGN KEY([class_id])
REFERENCES [dbo].[class] ([class_id])
GO
