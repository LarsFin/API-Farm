namespace ApiFarm.Utils
{
    /// <summary>
    /// Static definitions of messages for queries made.
    /// </summary>
    public struct QueryMessages
    {
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
        }
    }
}
