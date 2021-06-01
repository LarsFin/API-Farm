namespace ApiFarm.Utils
{
    /// <summary>
    /// Static definitions of messages for queries made.
    /// </summary>
    public struct ResponseMessages
    {
        /// <summary>
        /// Messages relating to JSON concerns.
        /// </summary>
        public struct JSON
        {
            /// <summary>
            /// Bad request message when JSON could not be serialized.
            /// </summary>
            public static string IsInvalid = "Invalid JSON in body.";
        }

        /// <summary>
        /// Messages relating to <see cref="Models.Impl.VideoGame"/> queries.
        /// </summary>
        public struct VideoGame
        {
            /// <summary>
            /// Bad request message when a <see cref="Models.Impl.VideoGame"/> is added but with a default DateReleased value.
            /// </summary>
            public static string RequiresDateReleased = "A date_released is required for a video game.";

            /// <summary>
            /// Bad request message when a <see cref="Models.Impl.VideoGame"/> is added but without a Name value.
            /// </summary>
            public static string RequiresName = "A name is required for a video game.";

            /// <summary>
            /// Bad request message when a <see cref="Models.Impl.VideoGame"/> is added but its date is not in the form 'dd/MM/yyyy'.
            /// </summary>
            /// <param name="invalidDate">The invalid date.</param>
            /// <returns>Message relating to an unsupported date format.</returns>
            public static string InvalidDateReleased(string invalidDate) => $"The provided date_released '{invalidDate}' is invalid.";

            /// <summary>
            /// Not found request message when a <see cref="Models.Impl.VideoGame"/> could not be found in storage.
            /// </summary>
            /// <param name="id">The identifier for which a <see cref="Models.Impl.VideoGame"/> could not be found.</param>
            /// <returns>Message relating to not found <see cref="Models.Impl.VideoGame"/>.</returns>
            public static string NotFound(uint id) => $"No video game with id '{id}' could be found.";
        }
    }
}
