--
-- PostgreSQL database dump
--

-- Dumped from database version 15.3 (Debian 15.3-1.pgdg110+1)
-- Dumped by pg_dump version 15.3 (Debian 15.3-1.pgdg110+1)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: accounts; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.accounts (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    username text,
    email text,
    password character varying(255) NOT NULL,
    email_verified boolean DEFAULT false,
    phone text
);


ALTER TABLE public.accounts OWNER TO root;

--
-- Name: accounts_id_seq; Type: SEQUENCE; Schema: public; Owner: root
--

CREATE SEQUENCE public.accounts_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.accounts_id_seq OWNER TO root;

--
-- Name: accounts_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: root
--

ALTER SEQUENCE public.accounts_id_seq OWNED BY public.accounts.id;


--
-- Name: bookings; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.bookings (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    user_id bigint,
    post_id bigint,
    start_date timestamp with time zone,
    end_date timestamp with time zone
);


ALTER TABLE public.bookings OWNER TO root;

--
-- Name: bookings_id_seq; Type: SEQUENCE; Schema: public; Owner: root
--

CREATE SEQUENCE public.bookings_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.bookings_id_seq OWNER TO root;

--
-- Name: bookings_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: root
--

ALTER SEQUENCE public.bookings_id_seq OWNED BY public.bookings.id;


--
-- Name: cats; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.cats (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name text,
    age bigint,
    text text,
    user_id bigint
);


ALTER TABLE public.cats OWNER TO root;

--
-- Name: cats_id_seq; Type: SEQUENCE; Schema: public; Owner: root
--

CREATE SEQUENCE public.cats_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.cats_id_seq OWNER TO root;

--
-- Name: cats_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: root
--

ALTER SEQUENCE public.cats_id_seq OWNED BY public.cats.id;


--
-- Name: forgot_passwords; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.forgot_passwords (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    token character varying(255) NOT NULL,
    account_id bigint NOT NULL,
    expired boolean DEFAULT false
);


ALTER TABLE public.forgot_passwords OWNER TO root;

--
-- Name: forgot_passwords_id_seq; Type: SEQUENCE; Schema: public; Owner: root
--

CREATE SEQUENCE public.forgot_passwords_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.forgot_passwords_id_seq OWNER TO root;

--
-- Name: forgot_passwords_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: root
--

ALTER SEQUENCE public.forgot_passwords_id_seq OWNED BY public.forgot_passwords.id;


--
-- Name: images; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.images (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    path text,
    post_id bigint,
    rent_id bigint,
    rent_detail_id bigint,
    main_post_id bigint
);


ALTER TABLE public.images OWNER TO root;

--
-- Name: images_id_seq; Type: SEQUENCE; Schema: public; Owner: root
--

CREATE SEQUENCE public.images_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.images_id_seq OWNER TO root;

--
-- Name: images_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: root
--

ALTER SEQUENCE public.images_id_seq OWNED BY public.images.id;


--
-- Name: locations; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.locations (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name text NOT NULL,
    description text,
    address text NOT NULL,
    latitude text,
    longitude text,
    user_id bigint,
    street_name text,
    post_code text,
    city text
);


ALTER TABLE public.locations OWNER TO root;

--
-- Name: locations_id_seq; Type: SEQUENCE; Schema: public; Owner: root
--

CREATE SEQUENCE public.locations_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.locations_id_seq OWNER TO root;

--
-- Name: locations_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: root
--

ALTER SEQUENCE public.locations_id_seq OWNED BY public.locations.id;


--
-- Name: posts; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.posts (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    brand text,
    brand_model text,
    vehicle_type text,
    year bigint,
    transmission text,
    fuel_type text,
    price_per_day bigint,
    price_per_week bigint,
    price_per_month bigint,
    discount bigint,
    units bigint,
    available boolean DEFAULT true,
    user_id bigint,
    location_id bigint,
    bookable boolean DEFAULT true,
    body_color text,
    license_plate text,
    price_per_day_after_discount bigint,
    price_per_week_after_discount bigint,
    price_per_month_after_discount bigint,
    discount_percentage bigint
);


ALTER TABLE public.posts OWNER TO root;

--
-- Name: posts_id_seq; Type: SEQUENCE; Schema: public; Owner: root
--

CREATE SEQUENCE public.posts_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.posts_id_seq OWNER TO root;

--
-- Name: posts_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: root
--

ALTER SEQUENCE public.posts_id_seq OWNED BY public.posts.id;


--
-- Name: rent_details; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.rent_details (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    license_plate text,
    pickup_date timestamp with time zone,
    return_date timestamp with time zone,
    decline_reason text,
    status text DEFAULT 'Accepted'::text,
    text text,
    rent_id bigint,
    is_paid boolean DEFAULT false,
    estimated_price bigint,
    estimated_saved_price bigint,
    rent_days bigint,
    estimated_final_price bigint,
    estimated_normal_price bigint
);


ALTER TABLE public.rent_details OWNER TO root;

--
-- Name: rent_details_id_seq; Type: SEQUENCE; Schema: public; Owner: root
--

CREATE SEQUENCE public.rent_details_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.rent_details_id_seq OWNER TO root;

--
-- Name: rent_details_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: root
--

ALTER SEQUENCE public.rent_details_id_seq OWNED BY public.rent_details.id;


--
-- Name: rents; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.rents (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    text text,
    user_id bigint,
    post_id bigint,
    start_date timestamp with time zone,
    end_date timestamp with time zone,
    pickup_date timestamp with time zone,
    return_date timestamp with time zone,
    license_plate text,
    status text DEFAULT 'ReadyToPickup'::text,
    payment_method text DEFAULT 'Paylater'::text,
    is_cancelled boolean DEFAULT false,
    cancel_reason text,
    discount_code text,
    readonly boolean DEFAULT false
);


ALTER TABLE public.rents OWNER TO root;

--
-- Name: rents_id_seq; Type: SEQUENCE; Schema: public; Owner: root
--

CREATE SEQUENCE public.rents_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.rents_id_seq OWNER TO root;

--
-- Name: rents_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: root
--

ALTER SEQUENCE public.rents_id_seq OWNED BY public.rents.id;


--
-- Name: reviews; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.reviews (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    stars bigint,
    text text,
    user_id bigint,
    post_id bigint,
    rent_id bigint
);


ALTER TABLE public.reviews OWNER TO root;

--
-- Name: reviews_id_seq; Type: SEQUENCE; Schema: public; Owner: root
--

CREATE SEQUENCE public.reviews_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.reviews_id_seq OWNER TO root;

--
-- Name: reviews_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: root
--

ALTER SEQUENCE public.reviews_id_seq OWNED BY public.reviews.id;


--
-- Name: transactions; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.transactions (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);


ALTER TABLE public.transactions OWNER TO root;

--
-- Name: transactions_id_seq; Type: SEQUENCE; Schema: public; Owner: root
--

CREATE SEQUENCE public.transactions_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.transactions_id_seq OWNER TO root;

--
-- Name: transactions_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: root
--

ALTER SEQUENCE public.transactions_id_seq OWNED BY public.transactions.id;


--
-- Name: users; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.users (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    public_username text,
    name text,
    phone integer,
    about text,
    gender text,
    role text DEFAULT 'Basic'::text,
    is_active text DEFAULT 'true'::text,
    account_id bigint
);


ALTER TABLE public.users OWNER TO root;

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: root
--

CREATE SEQUENCE public.users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.users_id_seq OWNER TO root;

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: root
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- Name: accounts id; Type: DEFAULT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.accounts ALTER COLUMN id SET DEFAULT nextval('public.accounts_id_seq'::regclass);


--
-- Name: bookings id; Type: DEFAULT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.bookings ALTER COLUMN id SET DEFAULT nextval('public.bookings_id_seq'::regclass);


--
-- Name: cats id; Type: DEFAULT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.cats ALTER COLUMN id SET DEFAULT nextval('public.cats_id_seq'::regclass);


--
-- Name: forgot_passwords id; Type: DEFAULT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.forgot_passwords ALTER COLUMN id SET DEFAULT nextval('public.forgot_passwords_id_seq'::regclass);


--
-- Name: images id; Type: DEFAULT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.images ALTER COLUMN id SET DEFAULT nextval('public.images_id_seq'::regclass);


--
-- Name: locations id; Type: DEFAULT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.locations ALTER COLUMN id SET DEFAULT nextval('public.locations_id_seq'::regclass);


--
-- Name: posts id; Type: DEFAULT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.posts ALTER COLUMN id SET DEFAULT nextval('public.posts_id_seq'::regclass);


--
-- Name: rent_details id; Type: DEFAULT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.rent_details ALTER COLUMN id SET DEFAULT nextval('public.rent_details_id_seq'::regclass);


--
-- Name: rents id; Type: DEFAULT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.rents ALTER COLUMN id SET DEFAULT nextval('public.rents_id_seq'::regclass);


--
-- Name: reviews id; Type: DEFAULT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.reviews ALTER COLUMN id SET DEFAULT nextval('public.reviews_id_seq'::regclass);


--
-- Name: transactions id; Type: DEFAULT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.transactions ALTER COLUMN id SET DEFAULT nextval('public.transactions_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- Data for Name: accounts; Type: TABLE DATA; Schema: public; Owner: root
--

COPY public.accounts (id, created_at, updated_at, deleted_at, username, email, password, email_verified, phone) FROM stdin;
1	2023-10-16 00:32:18.873911+00	2023-10-16 00:32:18.873911+00	\N	pinguin	pinguin@gmail.com	$2a$10$tIy/bcMwxKeMbbIxFDSj5exdrp8/Ape5rgud3wEv.okUEKa.bjXy.	f	\N
2	2023-10-16 00:33:09.154503+00	2023-10-16 12:51:48.074684+00	\N	admin	admin@gmail.com	$2a$10$dtYVsd3i1yS8PhmdoHlHdONONbN4BfOHk/oYYjKHw2MeIAdpThDne	f	+6281225972197
4	2023-10-26 09:15:59.68761+00	2023-10-26 09:15:59.68761+00	\N	kucingimut1	kucingimut1@gmail.com	$2a$10$jEAiBfTPScKqOIJ7fXBYfOFBDKAlCeS3ZPI0zpwfK7G5xoqor56Gm	f	
5	2023-10-29 08:30:20.117716+00	2023-10-29 08:30:20.117716+00	\N	kucingimut2	kucingimut2@gmail.com	$2a$10$eHf.tAN3JmASxB9rG0PxEOiGKXO9qU.PzAU7sSn4Zymln/yuGxfzi	f	
6	2023-11-11 14:08:54.128692+00	2023-11-11 14:08:54.128692+00	\N	kucing	kucing@gmail.com	$2a$10$QHGo/Wy30o.xeqiNVoW4PeOucHT/L0WAz/QFJw9I9v3OH8jHDAjJm	f	
7	2023-11-27 13:33:51.24825+00	2023-11-27 13:33:51.24825+00	\N	kambing	risqi.app.dev@gmail.com	$2a$10$M4k4.z99FBeluI08Y01yH.KKl5Vd.2KqF1OEQm1cjF1Q8yiNp4oiC	f	
\.


--
-- Data for Name: bookings; Type: TABLE DATA; Schema: public; Owner: root
--

COPY public.bookings (id, created_at, updated_at, deleted_at, user_id, post_id, start_date, end_date) FROM stdin;
\.


--
-- Data for Name: cats; Type: TABLE DATA; Schema: public; Owner: root
--

COPY public.cats (id, created_at, updated_at, deleted_at, name, age, text, user_id) FROM stdin;
2	2023-10-21 02:47:29.138281+00	2023-10-21 02:47:29.138281+00	\N	garfield 2	2		2
3	2023-10-21 02:48:17.291514+00	2023-10-21 02:48:17.291514+00	\N	garfield 2	0		2
4	2023-10-21 10:17:39.667682+00	2023-10-21 10:17:39.667682+00	\N	cutecat	2	very cute	2
1	2023-10-21 02:46:40.250896+00	2023-10-21 10:18:02.281424+00	\N	kucing imut	2	the big cat	2
5	2023-10-29 08:35:46.496585+00	2023-10-29 08:35:46.496585+00	\N	verycutecat	2	very cute	2
6	2023-10-29 12:51:56.919943+00	2023-10-29 12:51:56.919943+00	\N	blacky	2	kucing imut	2
7	2023-10-29 12:53:14.717375+00	2023-10-29 12:53:14.717375+00	\N	blackzzzz	2	kucingimut	2
\.


--
-- Data for Name: forgot_passwords; Type: TABLE DATA; Schema: public; Owner: root
--

COPY public.forgot_passwords (id, created_at, updated_at, deleted_at, token, account_id, expired) FROM stdin;
\.


--
-- Data for Name: images; Type: TABLE DATA; Schema: public; Owner: root
--

COPY public.images (id, created_at, updated_at, deleted_at, path, post_id, rent_id, rent_detail_id, main_post_id) FROM stdin;
49	2023-11-21 14:02:28.164877+00	2023-11-21 14:02:28.164877+00	\N	static\\images\\7c28b907-6167-4887-a77d-bdadb219d1dc_1700575348156199200.png	\N	\N	\N	45
50	2023-11-21 15:09:01.945323+00	2023-11-21 15:09:01.945323+00	\N	static\\images\\ddace034-1d33-45a1-9cd7-9e79324c7e8f_1700579341942664900.png	\N	\N	\N	47
51	2023-11-21 15:13:19.669141+00	2023-11-21 15:13:19.669141+00	\N	static\\images\\32fa575c-14dc-4eb5-8ccb-6022bd781058_1700579599664486000.png	\N	\N	\N	48
52	2023-11-21 15:18:29.746222+00	2023-11-21 15:18:29.746222+00	\N	static\\images\\003ea5d0-a06f-4b01-aa1d-56c599d7939a_1700579909744494700.png	\N	\N	\N	49
53	2023-11-23 13:18:03.25226+00	2023-11-23 13:18:03.25226+00	\N	static\\images\\726b460c-1219-46a7-ae65-ca190f6d2aee_1700745483244898800.png	\N	\N	\N	50
54	2023-11-23 14:09:12.43093+00	2023-11-23 14:09:12.43093+00	\N	static\\images\\bd611821-1c35-4e1c-80e7-d3148f3c2210_1700748552413025900.png	\N	\N	\N	51
55	2023-11-29 00:05:27.458557+00	2023-11-29 00:05:27.458557+00	\N	static\\images\\b4c84041-9683-4899-ac05-d07ac23976d2_1701216327451240000.png	\N	\N	\N	52
56	2023-11-29 00:06:26.721156+00	2023-11-29 00:06:26.721156+00	\N	static\\images\\d2971931-81be-43a7-9953-7d8c61ce50a5_1701216386718798300.png	\N	\N	\N	53
57	2023-11-29 00:09:09.360428+00	2023-11-29 00:09:09.360428+00	\N	static\\images\\29c60006-62e2-48c4-bd56-ff5cfa489b3a_1701216549358850000.png	\N	\N	\N	54
58	2023-11-29 00:11:26.666067+00	2023-11-29 00:11:26.666067+00	\N	static\\images\\dd53a16e-5fe7-4509-ae79-96fd08786d97_1701216686664982500.png	\N	\N	\N	55
59	2023-11-29 00:15:02.699834+00	2023-11-29 00:15:02.699834+00	\N	static\\images\\16adf16a-d55b-45d1-810f-4805ce80d0cc_1701216902697859200.png	\N	\N	\N	56
60	2023-11-29 00:19:35.458348+00	2023-11-29 00:19:35.458348+00	\N	static\\images\\40c83015-d1bb-42a9-8adf-16b653120562_1701217175456906900.png	\N	\N	\N	57
61	2023-11-29 00:23:01.228112+00	2023-11-29 00:23:01.228112+00	\N	static\\images\\2cc134b1-1e4a-49ab-8b46-00f68107896e_1701217381226655000.png	\N	\N	\N	58
62	2023-11-29 00:25:29.837853+00	2023-11-29 00:25:29.837853+00	\N	static\\images\\c8985dcc-5da9-42c1-a1e8-983a8d6d02f5_1701217529833633600.png	\N	\N	\N	59
46	2023-11-10 13:14:54.887716+00	2023-11-10 13:14:54.887716+00	\N	static\\images\\d1938307-bb64-4c22-a762-405633b7a1bd_1699622094867508000.jpg	\N	\N	\N	43
47	2023-11-10 13:14:54.970929+00	2023-11-10 13:14:54.970929+00	\N	static\\images\\ec00d28a-76fa-4d5a-961a-aa1e7512b19d_1699622094962821100.jpg	43	\N	\N	\N
48	2023-11-10 13:14:55.017945+00	2023-11-10 13:14:55.017945+00	\N	static\\images\\0468af28-8b1b-40ec-9334-ebd69f69a935_1699622095013469700.jpg	43	\N	\N	\N
\.


--
-- Data for Name: locations; Type: TABLE DATA; Schema: public; Owner: root
--

COPY public.locations (id, created_at, updated_at, deleted_at, name, description, address, latitude, longitude, user_id, street_name, post_code, city) FROM stdin;
1	2023-10-21 13:01:37.788724+00	2023-10-21 13:01:37.788724+00	\N	gerage 2 big	new gerage	yogyakarta city			2	\N	\N	\N
2	2023-10-21 13:01:55.9017+00	2023-10-21 13:01:55.9017+00	\N	gerage	good gerage	yogyakarta city			2	\N	\N	\N
3	2023-10-22 08:09:31.564494+00	2023-10-22 08:09:31.564494+00	\N	test gerage	test gerage	yogyakarta city			2	\N	\N	\N
4	2023-10-22 08:34:43.367252+00	2023-10-22 08:34:43.367252+00	\N	okay gerage	test gerage	yogyakarta city			2	\N	\N	\N
\.


--
-- Data for Name: posts; Type: TABLE DATA; Schema: public; Owner: root
--

COPY public.posts (id, created_at, updated_at, deleted_at, brand, brand_model, vehicle_type, year, transmission, fuel_type, price_per_day, price_per_week, price_per_month, discount, units, available, user_id, location_id, bookable, body_color, license_plate, price_per_day_after_discount, price_per_week_after_discount, price_per_month_after_discount, discount_percentage) FROM stdin;
52	2023-11-29 00:05:27.233751+00	2023-11-29 00:05:27.233751+00	\N	Toyota	Avanza	SUV	2023	automatic	gasoline	100	500	1700	\N	\N	t	2	2	t	white	AB 232 D	90	450	1530	10
53	2023-11-29 00:06:26.687059+00	2023-11-29 00:06:26.687059+00	\N	BMW	M4	Sedan	2023	automatic	gasoline	300	1000	2500	\N	\N	t	2	1	t	yellow	AB 23	285	950	2375	5
54	2023-11-29 00:09:09.347529+00	2023-11-29 00:09:09.347529+00	\N	Honda	Civic Type R	Sedan	2023	automatic	gasoline	250	750	2400	\N	\N	t	2	2	t	white	G 232	238	713	2280	5
55	2023-11-29 00:11:26.654739+00	2023-11-29 00:11:26.654739+00	\N	Honda	Honda Civic	Sedan	2023	automatic	gasoline	150	500	1700	\N	\N	t	2	1	t	white	G 23	143	475	1615	5
56	2023-11-29 00:15:02.679391+00	2023-11-29 00:15:02.679391+00	\N	Mercedes-Benz	C class	Sedan	2022	automatic	gasoline	150	500	1700	\N	\N	t	2	2	t	white	A 231	143	475	1615	5
57	2023-11-29 00:19:35.423111+00	2023-11-29 00:19:35.423111+00	\N	Toyota	Prius	Sedan	2023	automatic	electric	200	700	2400	\N	\N	t	2	1	t	white	A 111	190	665	2280	5
58	2023-11-29 00:23:01.2102+00	2023-11-29 00:23:01.2102+00	\N	Volvo	XC90	SUV	2023	automatic	gasoline	200	700	2400	\N	\N	t	2	2	t	white	A 1112	190	665	2280	5
59	2023-11-29 00:25:29.821751+00	2023-11-29 00:25:29.821751+00	\N	Nissan	Skyline GTR	Sport	2022	automatic	gasoline	300	800	2700	\N	\N	t	2	1	t	orange	A 54	285	760	2565	5
43	2023-11-10 13:14:54.340415+00	2023-11-30 12:30:32.833566+00	\N	Lamborghini	Gallardo	Sport	2016	manual	gasoline	100	600	2500	\N	\N	t	2	1	f	yellow	a 213 g	90	540	2250	10
\.


--
-- Data for Name: rent_details; Type: TABLE DATA; Schema: public; Owner: root
--

COPY public.rent_details (id, created_at, updated_at, deleted_at, license_plate, pickup_date, return_date, decline_reason, status, text, rent_id, is_paid, estimated_price, estimated_saved_price, rent_days, estimated_final_price, estimated_normal_price) FROM stdin;
17	2023-11-24 15:28:46.981963+00	2023-11-24 15:30:21.757947+00	\N		\N	\N	unit is not ready yet	Declined		17	f	\N	0	3	300	300
\.


--
-- Data for Name: rents; Type: TABLE DATA; Schema: public; Owner: root
--

COPY public.rents (id, created_at, updated_at, deleted_at, text, user_id, post_id, start_date, end_date, pickup_date, return_date, license_plate, status, payment_method, is_cancelled, cancel_reason, discount_code, readonly) FROM stdin;
17	2023-11-24 15:28:46.968373+00	2023-11-24 15:30:21.812903+00	\N	\N	6	43	2023-11-14 12:00:00+00	2023-11-17 12:00:00+00	\N	\N	\N	ReadyToPickup	Paylater	f			t
\.


--
-- Data for Name: reviews; Type: TABLE DATA; Schema: public; Owner: root
--

COPY public.reviews (id, created_at, updated_at, deleted_at, stars, text, user_id, post_id, rent_id) FROM stdin;
\.


--
-- Data for Name: transactions; Type: TABLE DATA; Schema: public; Owner: root
--

COPY public.transactions (id, created_at, updated_at, deleted_at) FROM stdin;
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: root
--

COPY public.users (id, created_at, updated_at, deleted_at, public_username, name, phone, about, gender, role, is_active, account_id) FROM stdin;
1	2023-10-16 00:32:18.901221+00	2023-10-16 00:32:18.901221+00	\N	mnovF8GEJjYDZPXXG2fZ		0			Basic	true	1
4	2023-10-26 09:15:59.694125+00	2023-10-26 09:15:59.694125+00	\N	8OMigep91Zfjm1NFqnG0		\N			Basic	true	4
5	2023-10-29 08:30:20.145514+00	2023-10-29 08:30:20.145514+00	\N	X6DYAPcUiFMMS0o4jboH		\N			Basic	true	5
6	2023-11-11 14:08:54.154693+00	2023-11-11 14:08:54.154693+00	\N	MKapEov1VIJ4f6Hm72U7		\N			Basic	true	6
7	2023-11-27 13:33:51.337763+00	2023-11-27 13:33:51.337763+00	\N	Ira38BtH9xG9hgflHbnq		\N			Basic	true	7
2	2023-10-16 00:33:09.160505+00	2023-11-30 12:17:39.977602+00	\N	admincat	The admin	0	I'm admin	Male	Admin	true	2
\.


--
-- Name: accounts_id_seq; Type: SEQUENCE SET; Schema: public; Owner: root
--

SELECT pg_catalog.setval('public.accounts_id_seq', 7, true);


--
-- Name: bookings_id_seq; Type: SEQUENCE SET; Schema: public; Owner: root
--

SELECT pg_catalog.setval('public.bookings_id_seq', 1, false);


--
-- Name: cats_id_seq; Type: SEQUENCE SET; Schema: public; Owner: root
--

SELECT pg_catalog.setval('public.cats_id_seq', 7, true);


--
-- Name: forgot_passwords_id_seq; Type: SEQUENCE SET; Schema: public; Owner: root
--

SELECT pg_catalog.setval('public.forgot_passwords_id_seq', 1, false);


--
-- Name: images_id_seq; Type: SEQUENCE SET; Schema: public; Owner: root
--

SELECT pg_catalog.setval('public.images_id_seq', 62, true);


--
-- Name: locations_id_seq; Type: SEQUENCE SET; Schema: public; Owner: root
--

SELECT pg_catalog.setval('public.locations_id_seq', 4, true);


--
-- Name: posts_id_seq; Type: SEQUENCE SET; Schema: public; Owner: root
--

SELECT pg_catalog.setval('public.posts_id_seq', 59, true);


--
-- Name: rent_details_id_seq; Type: SEQUENCE SET; Schema: public; Owner: root
--

SELECT pg_catalog.setval('public.rent_details_id_seq', 17, true);


--
-- Name: rents_id_seq; Type: SEQUENCE SET; Schema: public; Owner: root
--

SELECT pg_catalog.setval('public.rents_id_seq', 17, true);


--
-- Name: reviews_id_seq; Type: SEQUENCE SET; Schema: public; Owner: root
--

SELECT pg_catalog.setval('public.reviews_id_seq', 1, false);


--
-- Name: transactions_id_seq; Type: SEQUENCE SET; Schema: public; Owner: root
--

SELECT pg_catalog.setval('public.transactions_id_seq', 1, false);


--
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: root
--

SELECT pg_catalog.setval('public.users_id_seq', 7, true);


--
-- Name: accounts accounts_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.accounts
    ADD CONSTRAINT accounts_pkey PRIMARY KEY (id);


--
-- Name: bookings bookings_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.bookings
    ADD CONSTRAINT bookings_pkey PRIMARY KEY (id);


--
-- Name: cats cats_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.cats
    ADD CONSTRAINT cats_pkey PRIMARY KEY (id);


--
-- Name: forgot_passwords forgot_passwords_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.forgot_passwords
    ADD CONSTRAINT forgot_passwords_pkey PRIMARY KEY (id);


--
-- Name: images images_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.images
    ADD CONSTRAINT images_pkey PRIMARY KEY (id);


--
-- Name: locations locations_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.locations
    ADD CONSTRAINT locations_pkey PRIMARY KEY (id);


--
-- Name: posts posts_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.posts
    ADD CONSTRAINT posts_pkey PRIMARY KEY (id);


--
-- Name: rent_details rent_details_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.rent_details
    ADD CONSTRAINT rent_details_pkey PRIMARY KEY (id);


--
-- Name: rents rents_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.rents
    ADD CONSTRAINT rents_pkey PRIMARY KEY (id);


--
-- Name: reviews reviews_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.reviews
    ADD CONSTRAINT reviews_pkey PRIMARY KEY (id);


--
-- Name: transactions transactions_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.transactions
    ADD CONSTRAINT transactions_pkey PRIMARY KEY (id);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: idx_accounts_deleted_at; Type: INDEX; Schema: public; Owner: root
--

CREATE INDEX idx_accounts_deleted_at ON public.accounts USING btree (deleted_at);


--
-- Name: idx_accounts_email; Type: INDEX; Schema: public; Owner: root
--

CREATE UNIQUE INDEX idx_accounts_email ON public.accounts USING btree (email);


--
-- Name: idx_accounts_username; Type: INDEX; Schema: public; Owner: root
--

CREATE UNIQUE INDEX idx_accounts_username ON public.accounts USING btree (username);


--
-- Name: idx_bookings_deleted_at; Type: INDEX; Schema: public; Owner: root
--

CREATE INDEX idx_bookings_deleted_at ON public.bookings USING btree (deleted_at);


--
-- Name: idx_cats_deleted_at; Type: INDEX; Schema: public; Owner: root
--

CREATE INDEX idx_cats_deleted_at ON public.cats USING btree (deleted_at);


--
-- Name: idx_forgot_passwords_deleted_at; Type: INDEX; Schema: public; Owner: root
--

CREATE INDEX idx_forgot_passwords_deleted_at ON public.forgot_passwords USING btree (deleted_at);


--
-- Name: idx_images_deleted_at; Type: INDEX; Schema: public; Owner: root
--

CREATE INDEX idx_images_deleted_at ON public.images USING btree (deleted_at);


--
-- Name: idx_locations_deleted_at; Type: INDEX; Schema: public; Owner: root
--

CREATE INDEX idx_locations_deleted_at ON public.locations USING btree (deleted_at);


--
-- Name: idx_posts_deleted_at; Type: INDEX; Schema: public; Owner: root
--

CREATE INDEX idx_posts_deleted_at ON public.posts USING btree (deleted_at);


--
-- Name: idx_rent_details_deleted_at; Type: INDEX; Schema: public; Owner: root
--

CREATE INDEX idx_rent_details_deleted_at ON public.rent_details USING btree (deleted_at);


--
-- Name: idx_rents_deleted_at; Type: INDEX; Schema: public; Owner: root
--

CREATE INDEX idx_rents_deleted_at ON public.rents USING btree (deleted_at);


--
-- Name: idx_reviews_deleted_at; Type: INDEX; Schema: public; Owner: root
--

CREATE INDEX idx_reviews_deleted_at ON public.reviews USING btree (deleted_at);


--
-- Name: idx_transactions_deleted_at; Type: INDEX; Schema: public; Owner: root
--

CREATE INDEX idx_transactions_deleted_at ON public.transactions USING btree (deleted_at);


--
-- Name: idx_users_deleted_at; Type: INDEX; Schema: public; Owner: root
--

CREATE INDEX idx_users_deleted_at ON public.users USING btree (deleted_at);


--
-- Name: idx_users_public_username; Type: INDEX; Schema: public; Owner: root
--

CREATE UNIQUE INDEX idx_users_public_username ON public.users USING btree (public_username);


--
-- Name: images fk_posts_images; Type: FK CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.images
    ADD CONSTRAINT fk_posts_images FOREIGN KEY (post_id) REFERENCES public.posts(id);


--
-- Name: reviews fk_posts_reviews; Type: FK CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.reviews
    ADD CONSTRAINT fk_posts_reviews FOREIGN KEY (post_id) REFERENCES public.posts(id);


--
-- Name: images fk_rents_images; Type: FK CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.images
    ADD CONSTRAINT fk_rents_images FOREIGN KEY (rent_id) REFERENCES public.rents(id);


--
-- PostgreSQL database dump complete
--

