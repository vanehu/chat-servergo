/*
 *  Copyright (c) 2017, https://github.com/nebulaim
 *  All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package dataobject

type UserContactsDO struct {
	Id            int32  `db:"id"`
	OwnerUserId   int32  `db:"owner_user_id"`
	ContactUserId int32  `db:"contact_user_id"`
	IsBlocked     int8   `db:"is_blocked"`
	Date2         int32  `db:"date2"`
	IsDeleted     int8   `db:"is_deleted"`
	CreatedAt     string `db:"created_at"`
	UpdatedAt     string `db:"updated_at"`
}
